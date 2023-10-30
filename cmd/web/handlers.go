package main

import (
	"net/http"
	"strings"

	"github.com/LachlanStephan/mysite_server/cmd/files"
	"github.com/LachlanStephan/mysite_server/internal/validator"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	err := app.files.RenderFile(w, files.HomeTemplate)

	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) blog(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/blog" {
		app.notFound(w)
		return
	}

	err := app.files.RenderFile(w, files.BlogsTemplate)

	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) viewBlog(w http.ResponseWriter, r *http.Request) {
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
