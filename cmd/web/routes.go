package main
import (
	"github.com/bmizerany/pat" // New import
	"github.com/justinas/alice"
	"net/http"
)
func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Post("/comment/create", http.HandlerFunc(app.createComment))
	mux.Post("/delete", http.HandlerFunc(app.deleteComment))// Moved down
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)
}

