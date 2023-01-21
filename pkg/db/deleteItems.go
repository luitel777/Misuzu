package db

import (
	"fmt"
	"log"
)

func DeleteMessageDB() {
	delete_db := MisuzuDB.InitMisuzuDatabase(MisuzuDB{})

	fmt.Println("Deleting db")
	delete_db.Delete(&[]MisuzuModel{}, "id > 0")
	sqlDB, err := delete_db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()
}
