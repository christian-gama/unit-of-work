package user

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"not null;unique;size:255" binding:"required" json:"name"`
	Age  int    `gorm:"not null;size:3" binding:"required" json:"age"`
}

func Migrate(db *gorm.DB) {
	log.Println("Migrating the User table")
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Error while migrating the User table: %v", err)
	}
}
