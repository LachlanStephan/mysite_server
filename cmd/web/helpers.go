package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Print(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request, customError error) {
	data := app.newFileData(r)
	data.CustomError = customError
	template, err := app.getFile(notFoundTemplate)
	if err != nil {
		app.serverError(w, err)
	}

	app.render(w, http.StatusNotFound, template, data)
}

func (app *application) pathExists(path string, r *http.Request) bool {
	return r.URL.Path == path
}

func (app *application) pathContains(path string, r *http.Request) bool {
	return strings.Contains(r.URL.Path, path)
}
