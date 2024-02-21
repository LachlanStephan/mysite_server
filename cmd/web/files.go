package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type FileData struct {
	CurrentYear int
	BlogLinks   []*BlogLinks
	CustomError error
}

// creating an object to store helper functions that the templates may use
// these must return only 1 value || 1 value && err
var functions = template.FuncMap{}

// create in memory cache to store templates and prevent repeating ourselves in the handler funcs
func newFileCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	blogs, err := getFilesFromPath(blogsPath)
	if err != nil {
		return nil, err
	}

	pages, err := getFilesFromPath(pagesPath)
	if err != nil {
		return nil, err
	}

	pages = append(pages, blogs...)

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func getFilesFromPath(path string) ([]string, error) {
	blogs, err := filepath.Glob(path)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (app *application) newFileData(r *http.Request) *FileData {
	return &FileData{
		CurrentYear: time.Now().Year(),
		CustomError: nil,
	}
}

func (app *application) render(w http.ResponseWriter, status int, t *template.Template, data *FileData) {
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
		err := fmt.Errorf("looks like %s does not exist", file)
		return nil, err
	}

	return ts, nil
}

// can be more than one file ext eg. ".tmpl.html"
func stripFileExt(name string) string {
	r := name
	extCount := strings.Count(name, ".")
	for i := 0; i < extCount; i++ {
		r = strings.TrimSuffix(r, filepath.Ext(r))
	}
	return r
}
