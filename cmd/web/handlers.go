package main

import (
	"fmt"
	"net/http"
	"strconv"
)
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	s, err := app.comments.All()
	if err != nil {
		fmt.Println("here")
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.tmpl", &templateData{
		Comments: s,
	})

}

func (app *application) createComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	author := r.PostForm.Get("author")
	comment := r.PostForm.Get("comment")
	pID, parentErr := strconv.Atoi(r.PostForm.Get("parentId"))
	if parentErr != nil {
		app.serverError(w, err)
		return
	}
	id, err := app.comments.Insert(author, comment, pID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write([]byte(strconv.Itoa(id)))

}
func (app *application)  deleteComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(r.PostForm.Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	err = app.comments.Delete(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

}