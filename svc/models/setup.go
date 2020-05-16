package models

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// InitDB sets up db and runs migrations
func InitDB() *gorm.DB {
	db, err := gorm.Open("postgres", createDbString())
	if err != nil {
		panic("Failed to connect to database!")
	}

	runMigrations(db)

	return db
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&Todo{})
}

func createDbString() string {
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")
	name := os.Getenv("DBNAME")

	return "host=" + host + " port=" + port + " user=" + user + " dbname=" + name + " password=" + pass + " sslmode=disable"
}
