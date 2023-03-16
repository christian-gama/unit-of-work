package user

import (
	"log"

	"gorm.io/gorm"
)

// User is a struct representing a User in the system.
// It embeds GORM's Model struct and adds Name and Age fields.
type User struct {
	gorm.Model
	Name string `gorm:"not null;unique;size:255" binding:"required" json:"name"`
	Age  int    `gorm:"not null;size:3" binding:"required" json:"age"`
}

// Migrate is a function that creates the User table in the database.
// It takes a GORM database instance as an argument and executes the necessary migration to create the table.
func Migrate(db *gorm.DB) {
	log.Println("Migrating the User table")
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Error while migrating the User table: %v", err)
	}
}
