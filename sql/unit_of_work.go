package sql

import (
	"log"

	"github.com/christian-gama/uow/uow"
	"gorm.io/gorm"
)

type UnitOfWork struct {
	db *gorm.DB
	Tx *gorm.DB
}

func (uow *UnitOfWork) Commit() error {
	log.Println("Commiting the transaction")

	err := uow.Tx.Commit().Error
	uow.Tx = uow.db

	return err
}

func (uow *UnitOfWork) Rollback() error {
	log.Println("Rolling back the transaction")

	err := uow.Tx.Rollback().Error
	uow.Tx = uow.db

	return err
}

func (uow *UnitOfWork) Begin() error {
	log.Println("Starting a new transaction")

	uow.Tx = uow.db.Begin()

	return uow.Tx.Error
}

func NewUnitOfWork(db *gorm.DB) uow.UnitOfWork {
	return &UnitOfWork{db: db, Tx: db}
}
