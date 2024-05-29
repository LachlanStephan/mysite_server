package main

import (
	"net/http"
)

func runServer(port string, app *application) {
	mux := http.NewServeMux()
	routes := app.getRoutes()

	p, err := getPathPrefix()
	if err != nil {
		app.errorLog.Fatal(err)
	}

	fileServer := http.FileServer(http.Dir(p + "/ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	for k, v := range routes {
		mux.HandleFunc(k, v)
	}

	app.infoLog.Printf("Starting on port %s", port)
	err = http.ListenAndServe(":"+port, mux)
	app.errorLog.Fatal(err)
}

func (app *application) getRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	list := map[string]func(http.ResponseWriter, *http.Request){
		"/":            app.home,
		"/blogs":       app.blogs,
		"/books":       app.books,
		"/blogs/view/": app.viewContent,
		"/books/view/": app.viewContent,
		"/favicon.ico": app.favicon,
	}

	return list
}
