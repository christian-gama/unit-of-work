package sql

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Conn is a function that connects to a PostgreSQL database using the provided connection details and returns a GORM database instance.
func Conn() *gorm.DB {
	log.Printf("Connecting to database at %s:%d", dsn.Host, dsn.Port)

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}

	return db
}
