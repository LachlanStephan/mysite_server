package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	httpMethod string
	path       string
	handler    func(http.ResponseWriter, *http.Request)
}

func runServer(port string, app *application) {
	router := httprouter.New()
	routes := app.getRoutes()

	for i := 0; i < len(routes); i++ {
		route := routes[i]
		router.HandlerFunc(route.httpMethod, route.path, route.handler)
	}

	p, err := getPathPrefix()
	if err != nil {
		app.errorLog.Fatal(err)
	}

	fileServer := http.FileServer(http.Dir(p + "/ui/static"))
	router.Handler(http.MethodGet, "/static/", http.StripPrefix("/static", fileServer))

	app.infoLog.Printf("Starting on port %s", port)
	err = http.ListenAndServe(":"+port, app.rateLimit(router))
	app.errorLog.Fatal(err)
}

func (app *application) getRoutes() []Route {
	routes := []Route{
		Route{
			httpMethod: http.MethodGet,
			path:       "/",
			handler:    app.home,
		},
		Route{
			httpMethod: http.MethodGet,
			path:       "/blogs/view/:id",
			handler:    app.viewContent,
		},
		Route{
			httpMethod: http.MethodGet,
			path:       "/favicon.ico",
			handler:    app.favicon,
		},
	}

	return routes
}
