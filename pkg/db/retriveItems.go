package db

import "fmt"

func RetriveMessageDB() []MisuzuModel {
	retrive_db := MisuzuDB.InitMisuzuDatabase(MisuzuDB{})

	fmt.Println("Retriving db")
	models := []MisuzuModel{}
	retrive_db.Find(&[]MisuzuModel{}).Scan(&models)
	fmt.Println(models)

	return models
}
