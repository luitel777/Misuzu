package handlers

import (
	"fmt"
	"net/http"

	"github.com/luitel777/misuzu/pkg/db"
)

func GetMessage(w http.ResponseWriter, r *http.Request) {
	// var misuzu_time string

	// month := time.Now().Month()
	// day := time.Now().Day()
	// year := time.Now().Year()

	// misuzu_time += strconv.Itoa(day)
	// misuzu_time += month.String()
	// misuzu_time += strconv.Itoa(year)

	misuzu_message := r.FormValue("message")

	fmt.Println(misuzu_message)
	// fmt.Println(misuzu_time)

	db.SaveMessageDB(misuzu_message)
}
