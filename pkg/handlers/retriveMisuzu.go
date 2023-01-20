package handlers

import (
	"github.com/luitel777/misuzu/pkg/db"
)

func RetriveMessage(many int) []db.MisuzuModel {
	contents := db.RetriveMessageDB(many)
	return contents
}
