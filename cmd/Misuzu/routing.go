package main

import (
	"net/http"

	"github.com/luitel777/misuzu/pkg/handlers"
)

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		renderPage(w, r, "web/template/homePage.html", 10)
	case "/update":
		handlers.SaveMessage(w, r)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	case "/archive":
		renderPage(w, r, "web/template/archives.html", 2147483647)

	default:
		w.WriteHeader(http.StatusNotFound)
		renderErrorPage(w, r)
	}
}
