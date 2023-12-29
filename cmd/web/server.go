package main

import (
	"net/http"
)

func runServer(port string, app *application) error {
	mux := http.NewServeMux()
	routes := app.getRoutes()

	for k, v := range routes {
		mux.HandleFunc(k, v)
	}

	app.infoLog.Printf("Starting on port %s", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) getRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	list := map[string]func(http.ResponseWriter, *http.Request){
		"/":           app.home,
		"/blog":       app.blog,
		"/blog/view/": app.viewBlog,
	}

	return list
}
