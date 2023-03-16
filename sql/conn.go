package sql

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = &Dsn{
	Host:     "localhost",
	Port:     5432,
	User:     "postgres",
	Password: "postgres",
	DbName:   "postgres",
	SSLMode:  "disable",
}

func Conn() *gorm.DB {
	log.Printf("Connecting to database at %s:%d", dsn.Host, dsn.Port)

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}

	return db
}
