package main

import (
	"net/http"
	"strings"

	"github.com/LachlanStephan/mysite_server/internal/validator"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if !app.pathExists("/", r) {
		app.notFound(w, r)
		return
	}

	data := app.newFileData(r)

	app.render(w, http.StatusOK, homeTemplate, data)
}

func (app *application) blog(w http.ResponseWriter, r *http.Request) {
	if !app.pathExists("/blog", r) {
		app.notFound(w, r)
		return
	}

	data := app.newFileData(r)

	app.render(w, http.StatusOK, blogsTemplate, data)
}

func (app *application) viewBlog(w http.ResponseWriter, r *http.Request) {
	if !app.pathContains("/blog/view/", r) {
		app.notFound(w, r)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	length := len(parts) - 1
	id := parts[length]

	err := validator.InvalidId(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Output(0, id)
}
