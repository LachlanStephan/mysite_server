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

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	app.infoLog.Printf("Starting on port %s", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) getRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	list := map[string]func(http.ResponseWriter, *http.Request){
		"/":            app.home,
		"/blogs":       app.blogs,
		"/books":       app.books,
		"/blogs/view/": app.viewContent,
		"/books/view/": app.viewContent,
	}

	return list
}
