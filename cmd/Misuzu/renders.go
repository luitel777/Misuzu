package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type config struct {
	PageTitle string
	Content   template.HTML
}

func renderPage(w http.ResponseWriter, r *http.Request, path string) {
	con, err := os.ReadFile("dump")
	if err != nil {
		log.Println(err)
	}
	data := config{
		PageTitle: "Miso soup",
		Content:   template.HTML(con),
	}

	tmpl, err := template.ParseFiles(path)

	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, data)
}

func renderErrorPage(w http.ResponseWriter, r *http.Request) {
	data := config{
		PageTitle: "404 not found",
	}

	tmpl, err := template.ParseFiles("web/template/errorPage.html")

	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, data)
}
