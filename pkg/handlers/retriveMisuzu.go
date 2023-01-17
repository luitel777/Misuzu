package handlers

import (
	"github.com/luitel777/misuzu/pkg/db"
)

func RetriveMessage() []db.MisuzuModel {
	contents := db.RetriveMessageDB()
	return contents
}
