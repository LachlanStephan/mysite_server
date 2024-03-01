package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if !app.pathExists("/", r) {
		app.notFound(w, r, nil)
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

func (app *application) books(w http.ResponseWriter, r *http.Request) {
	if !app.pathExists("/books", r) {
		app.notFound(w, r, nil)
		return
	}

	f, err := app.getFile(booksTemplate)
	if err != nil {
		app.serverError(w, err)
		return
	}

	bookLinks, err := getContentLinks(booksPath, booksHrefPrefix)
	if err != nil {
		app.serverError(w, err)
	}

	data := app.newFileData(r)
	data.ContentLinks = bookLinks

	app.render(w, http.StatusOK, f, data)
}

func (app *application) blogs(w http.ResponseWriter, r *http.Request) {
	if !app.pathExists("/blogs", r) {
		app.notFound(w, r, nil)
		return
	}

	f, err := app.getFile(blogsTemplate)
	if err != nil {
		app.serverError(w, err)
		return
	}

	blogLinks, err := getContentLinks(blogsPath, blogHrefPrefix)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newFileData(r)
	data.ContentLinks = blogLinks

	app.render(w, http.StatusOK, f, data)
}

func (app *application) viewContent(w http.ResponseWriter, r *http.Request) {
	if !app.pathContains("/blogs/view/", r) && !app.pathContains("books/view/", r) {
		app.notFound(w, r, nil)
		return
	}

	fileName := getFileNameFromPath(r.URL.Path)

	f, err := app.getFile(fileName)
	if err != nil {
		app.notFound(w, r, err)
		return
	}

	data := app.newFileData(r)

	app.render(w, http.StatusOK, f, data)
}
