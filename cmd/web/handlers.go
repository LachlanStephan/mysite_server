package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if !app.isCorrectPath("/", r) {
		app.notFound(w, r, nil)
		return
	}

	f, err := app.getFile(homeTemplate)
	if err != nil {
		app.serverError(w, err)
		return
	}

	links, err := getContentLinks(blogsPath, blogHrefPrefix, true)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newFileData()
	data.ContentLinks = links

	app.render(w, http.StatusOK, f, data)
}

func (app *application) viewContent(w http.ResponseWriter, r *http.Request) {
	if !app.pathContains("/view/", r) {
		app.notFound(w, r, nil)
		return
	}

	fileName := getFileNameFromPath(r.URL.Path)

	f, err := app.getFile(fileName)
	if err != nil {
		app.notFound(w, r, err)
		return
	}

	links, err := getContentLinks(blogsPath, blogHrefPrefix, false)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newFileData()
	data.ContentLinks = links

	app.render(w, http.StatusOK, f, data)
}

func (app *application) favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/static/img/cartoon-me.ico")
}
