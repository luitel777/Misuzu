package handlers

import (
	"net/http"

	"github.com/luitel777/misuzu/pkg/db"
)

func SaveMessage(w http.ResponseWriter, r *http.Request) {
	misuzu_message := r.FormValue("message")
	misuzu_title := r.FormValue("title")

	db.SaveMessageDB(misuzu_message, misuzu_title)
}
