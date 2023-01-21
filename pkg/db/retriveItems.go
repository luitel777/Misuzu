package db

import (
	"fmt"
	"log"
)

func RetriveMessageDB(many int) []MisuzuModel {
	retrive_db := MisuzuDB.InitMisuzuDatabase(MisuzuDB{})

	fmt.Println("Retriving db")
	models := []MisuzuModel{}
	retrive_db.Last(&[]MisuzuModel{}).Limit(many).Scan(&models)
	fmt.Println(models)
	sqlDB, err := retrive_db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()

	return models
}
