package db

import (
	"html/template"
	"log"
	"time"
)

func SaveMessageDB(message string) {
	updatemodel := MisuzuModel{
		Date:       time.Now().Unix(),
		Title:      " ",
		Content:    template.HTML(message),
		References: "https://misuzu.com",
	}
	update_db := MisuzuDB.InitMisuzuDatabase(MisuzuDB{})

	update_db.AutoMigrate(&MisuzuModel{})
	update_db.Create(&updatemodel)

	sqlDB, err := update_db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()
}
