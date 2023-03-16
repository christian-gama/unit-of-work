package sql

import (
	"github.com/christian-gama/uow/uow"
	"gorm.io/gorm"
)

// unitOfWorkImpl is a struct that implements the unitOfWorkImpl interface and holds a GORM database connection.
type unitOfWorkImpl struct {
	db            *gorm.DB
	isTransaction bool
}

// Commit is a function that commits the current transaction. If there is no transaction, it does nothing.
func (u *unitOfWorkImpl) Commit() error {
	if u.isTransaction {
		return u.db.Commit().Error
	}

	return nil
}

// Rollback is a function that rolls back the current transaction. If there is no transaction, it does nothing.
func (u *unitOfWorkImpl) Rollback() error {
	if u.isTransaction {
		return u.db.Rollback().Error
	}

	return nil
}

// Transaction is a function that executes the provided function in a transaction.
// If the unit of work is already in a transaction, it executes the provided function without creating a new transaction.
func (u *unitOfWorkImpl) Transaction(fn func(uow uow.UnitOfWork) error) error {
	if u.isTransaction {
		return fn(u)
	}

	return u.db.Transaction(func(tx *gorm.DB) error {
		return fn(&unitOfWorkImpl{db: tx, isTransaction: true})
	})
}

// NewUnitOfWork is a function that creates a new unit of work instance using the provided GORM database connection.
func NewUnitOfWork(db *gorm.DB) *unitOfWorkImpl {
	return &unitOfWorkImpl{db: db}
}
