package main

import (
	"fmt"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("web"))
	http.Handle("/web/", http.StripPrefix("/web", fs))

	http.HandleFunc("/", pathHandler)

	fmt.Println("listening on 9090")

	http.ListenAndServe(":9090", nil)
}
