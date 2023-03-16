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

// Commit is a function that commits the transaction in progress.
// It sets the transaction to the parent database instance.
// It returns any errors encountered during the operation.
func (uow *UnitOfWork) Commit() error {
	log.Println("Commiting the transaction")

	err := uow.Tx.Commit().Error
	uow.Tx = uow.db

	return err
}

// Rollback is a function that rolls back the transaction in progress.
// It sets the transaction to the parent database instance.
// It returns any errors encountered during the operation.
func (uow *UnitOfWork) Rollback() error {
	log.Println("Rolling back the transaction")

	err := uow.Tx.Rollback().Error
	uow.Tx = uow.db

	return err
}

// Begin is a function that starts a new transaction.
// It sets the transaction to the parent database instance.
// It returns any errors encountered during the operation.
func (uow *UnitOfWork) Begin() error {
	log.Println("Starting a new transaction")

	uow.Tx = uow.db.Begin()

	return uow.Tx.Error
}

func NewUnitOfWork(db *gorm.DB) uow.UnitOfWork {
	return &UnitOfWork{db: db, Tx: db}
}
