package sql

import (
	"context"

	"github.com/christian-gama/uow/uow"
	"gorm.io/gorm"
)

type Repository interface {
	DB(ctx context.Context, uow uow.UnitOfWork) *gorm.DB
}

type repositoryImpl struct {
	db *gorm.DB
}

// DB is a function that returns a GORM database connection from the provided unit of work.
// If the unit of work is not of type UnitOfWork or it's nil, it returns a new GORM database connection.
func (r *repositoryImpl) DB(ctx context.Context, u uow.UnitOfWork) *gorm.DB {
	gormUnitOfWork, ok := u.(*unitOfWorkImpl)
	if !ok {
		return r.db.Session(&gorm.Session{NewDB: true}).WithContext(ctx)
	}

	return gormUnitOfWork.db.Session(&gorm.Session{NewDB: true}).WithContext(ctx)
}

// NewRepository is a function that creates a new user repository instance using the provided SQL repository.
func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db}
}
