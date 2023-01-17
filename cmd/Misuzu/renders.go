package main

import (
	"html/template"
	"net/http"

	"github.com/luitel777/misuzu/pkg/handlers"
)

type config struct {
	PageTitle string
	Content   template.HTML
	Date      int64
}

func renderPage(w http.ResponseWriter, r *http.Request, path string) {

	misuzu_contents := handlers.RetriveMessage()

	data := make([]config, 10)

	// data[0] = config{
	// 	PageTitle: "Miso soup",
	// 	Content:   template.HTML("ooo"),
	// }

	for i, k := range misuzu_contents {
		data[i] = config{
			PageTitle: "Miso soup",
			Date:      k.Date,
			Content:   template.HTML(k.Content),
		}
	}

	tmpl, err := template.ParseFiles(path)

	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, data[1])
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
