package db

import (
	"fmt"
	"log"
)

func RetriveMessageDB() []MisuzuModel {
	retrive_db := MisuzuDB.InitMisuzuDatabase(MisuzuDB{})

	fmt.Println("Retriving db")
	models := []MisuzuModel{}
	retrive_db.Find(&[]MisuzuModel{}).Scan(&models)
	fmt.Println(models)
	sqlDB, err := retrive_db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()

	return models
}
