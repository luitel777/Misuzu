package db

import "time"

func SaveMessageDB(message string) {
	updatemodel := MisuzuModel{
		Date:       time.Month(time.Now().Local().Unix()),
		Title:      " ",
		Content:    message,
		References: "https://misuzu.com",
	}
	update_db := MisuzuDB.InitMisuzuDatabase(MisuzuDB{})

	update_db.AutoMigrate(&MisuzuModel{})
	update_db.Create(&updatemodel)
}
