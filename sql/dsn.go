package sql

import "fmt"

// Dsn is a struct that represents a database connection string.
type Dsn struct {
	Host     string
	Port     uint
	User     string
	Password string
	DbName   string
	SSLMode  string
}

// String is a function that returns a string representation of the Dsn struct.
func (d *Dsn) String() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		d.Host,
		d.User,
		d.Password,
		d.DbName,
		d.Port,
		d.SSLMode,
	)
}

var dsn = &Dsn{
	Host:     "db",
	Port:     5432,
	User:     "postgres",
	Password: "postgres",
	DbName:   "postgres",
	SSLMode:  "disable",
}
