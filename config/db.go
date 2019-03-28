package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./sqlite.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//db.AutoMigrate(&Person{})

	return db
}
