package sql

import "fmt"

type Dsn struct {
	Host     string
	Port     uint
	User     string
	Password string
	DbName   string
	SSLMode  string
}

func (d *Dsn) String() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		d.Host,
		d.User,
		d.Password,
		d.DbName,
		d.Port,
		d.SSLMode)
}
