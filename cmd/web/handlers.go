package main

import (
	"net/http"
	"strings"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if !app.pathExists("/", r) {
		app.notFound(w, r)
		return
	}

	f, err := app.getFile(homeTemplate)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newFileData(r)

	app.render(w, http.StatusOK, f, data)
}

func (app *application) blog(w http.ResponseWriter, r *http.Request) {
	if !app.pathExists("/blog", r) {
		app.notFound(w, r)
		return
	}

	data := app.newFileData(r)

	f, err := app.getFile(blogsTemplate)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, http.StatusOK, f, data)
}

func (app *application) viewBlog(w http.ResponseWriter, r *http.Request) {
	if !app.pathContains("/blog/view/", r) {
		app.notFound(w, r)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	length := len(parts) - 1
	id := parts[length]
	id = id + ".tmpl.html"

	f, err := app.getFile(id)
	if err != nil {
		app.notFound(w, r)
		return
	}

	data := app.newFileData(r)

	app.render(w, http.StatusOK, f, data)
}
