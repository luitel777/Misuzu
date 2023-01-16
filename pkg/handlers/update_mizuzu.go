package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Update(w http.ResponseWriter, r *http.Request) {
	con, err := os.OpenFile("dump", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	var message string

	month := time.Now().Month()
	day := time.Now().Day()
	year := time.Now().Year()

	fmt.Println(strconv.Itoa(day), month, year)

	message += strconv.Itoa(day)
	message += month.String()
	message += strconv.Itoa(year)
	message += ": "
	message += r.FormValue("message")
	message += "<br>"
	fmt.Println(message)

	_, err = con.Write([]byte(message))
	if err != nil {
		log.Println("Openfile", err)
	}
}
