package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/luitel777/misuzu/pkg/db"
	"github.com/luitel777/misuzu/pkg/handlers"
)

type homepageValues struct {
	PageTitle string
	Misuzu    []db.MisuzuModel
}

func renderPage(w http.ResponseWriter, r *http.Request, path string, many int) {

	misuzu_contents := handlers.RetriveMessage(many)

	data := homepageValues{}

	misuzuArray := []db.MisuzuModel{}
	for i := range misuzu_contents {
		misuzuArray = append(misuzuArray, db.MisuzuModel{Date: misuzu_contents[i].Date, Content: misuzu_contents[i].Content, Title: misuzu_contents[i].Title})
	}

	data = homepageValues{
		PageTitle: "Miso Soup",
		Misuzu:    append([]db.MisuzuModel{}, misuzuArray...),
	}

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println("Retriving data")
	fmt.Println(misuzuArray)

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
}

func renderErrorPage(w http.ResponseWriter, r *http.Request) {
	data := homepageValues{
		PageTitle: "404 not found",
	}

	tmpl, err := template.ParseFiles("web/template/errorPage.html")

	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, data)
}
