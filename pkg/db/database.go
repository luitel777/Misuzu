package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MisuzuModel struct {
	Id         int `gorm:"primaryKey"`
	Date       time.Month
	Title      string
	Content    string
	References string
}

type MisuzuDB struct {
	MisuzuInstance *gorm.DB
}

func (MisuzuDB) InitMisuzuDatabase() *gorm.DB {

	dbURL := "postgres://pg:misuzu_password@localhost:5432/misuzu_db"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	return db
}

// func ping_database() {
// 	db := MisuzuDB.InitMisuzuDatabase(MisuzuDB{})

// 	testmodel := MisuzuModel{
// 		Date:  time.Now().Month(),
// 		Title: "Making a coffee",
// 		Content: `1) Add 1 tsb of sugar
// 				  2) Add 1 cup of water
// 				  3) Add 1 tsb of coffee
// 				  4) Boil it
// 				`,
// 		References: "https://something.com",
// 	}

// 	db.AutoMigrate(&MisuzuModel{})

// 	db.Create(&testmodel)

// 	db.Delete(&testmodel)

// }
