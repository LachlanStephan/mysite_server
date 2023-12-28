package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Print(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) render(w http.ResponseWriter, status int, page string, data *Files) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("looks like this template does not exist - %s", page)
		app.serverError(w, err)
		return
	}

	// we are going to check if the template is correct and safe to build
	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func (app *application) newTemplateData(r *http.Request) *Files {
	return &Files{
		CurrentYear: time.Now().Year(),
	}
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, http.StatusNotFound, notFoundTemplate, data)
}

func (app *application) pathExists(path string, r *http.Request) bool {
	return r.URL.Path == path
}

func (app *application) pathContains(path string, r *http.Request) bool {
	return strings.Contains(r.URL.Path, path)
}
