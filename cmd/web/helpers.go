package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"
	"strings"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Print(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) render(w http.ResponseWriter, status int, t *template.Template, data *Files) {
	// we are going to check if the template is correct and safe to build
	buf := new(bytes.Buffer)

	err := t.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func (app *application) getFile(file string) (*template.Template, error) {
	ts, ok := app.fileCache[file]
	if !ok {
		err := fmt.Errorf("looks like this file does not exist - %s", file)
		return nil, err
	}

	return ts, nil
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	data := app.newFileData(r)
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
