package main

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func (app *application) rateLimit(next http.Handler) http.Handler {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		go func() {
			time.Sleep(time.Minute)
			mu.Lock()

			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}

			mu.Unlock()
		}()

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			app.serverError(w, err)
			return
		}

		mu.Lock()

		if _, exists := clients[ip]; !exists {
			clients[ip] = &client{limiter: rate.NewLimiter(2, 4)}
		}
		clients[ip].lastSeen = time.Now()

		if !clients[ip].limiter.Allow() {
			mu.Unlock()
			app.rateLimitExceeded(w, r)
			return
		}

		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}
