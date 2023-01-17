package main

import (
	"net/http"

	"github.com/luitel777/misuzu/pkg/handlers"
)

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		renderPage(w, r, "web/template/homePage.html")

	case "/update":
		handlers.SaveMessage(w, r)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)

	default:
		w.WriteHeader(http.StatusNotFound)
		renderErrorPage(w, r)
	}
}
