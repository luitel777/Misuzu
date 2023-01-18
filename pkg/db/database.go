package db

import (
	"fmt"
	"html/template"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MisuzuModel struct {
	Id         int `gorm:"primaryKey"`
	Date       int64
	Title      string
	Content    template.HTML
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
