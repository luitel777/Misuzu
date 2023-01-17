package handlers

import (
	"fmt"
	"net/http"

	"github.com/luitel777/misuzu/pkg/db"
)

func SaveMessage(w http.ResponseWriter, r *http.Request) {
	misuzu_message := r.FormValue("message")

	fmt.Println(misuzu_message)

	db.SaveMessageDB(misuzu_message)
}
