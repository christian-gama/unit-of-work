package user

import (
	"log"
	"time"

	"gorm.io/gorm"
)

// User is a struct representing a User in the system.
// It embeds GORM's Model struct and adds Name, Age and Money fields.
type User struct {
	ID        uint           `gorm:"primarykey"`
	Age       int            `gorm:"not null;size:3" json:"age"`
	Money     int            `gorm:"not null;size:9" json:"money"`
	Name      string         `gorm:"not null;unique;size:255" json:"name"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Migrate is a function that creates the User table in the database.
// It takes a GORM database instance as an argument and executes the necessary migration to create the table.
func Migrate(db *gorm.DB) {
	log.Println("Migrating the User table")
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Error while migrating the User table: %v", err)
	}
}
