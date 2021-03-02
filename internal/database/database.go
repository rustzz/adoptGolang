package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func ConnectDatabase() *gorm.DB {
	var db, err = gorm.Open(sqlite.Open("db"), &gorm.Config{})
	if err != nil { log.Fatalln(err) }
	return db
}

func Migrate() {
	db := ConnectDatabase()
	//if err := db.AutoMigrate(); err != nil { return }
	dbConn, err := db.DB()
	if err != nil { return }
	if err = dbConn.Close(); err != nil { return }
	return
}
