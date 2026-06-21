package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
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
	params := httprouter.ParamsFromContext(r.Context())
	fileName := params.ByName("id") + htmlSuffix

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
