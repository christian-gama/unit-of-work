package sql

import (
	"log"

	"github.com/christian-gama/uow/uow"
	"gorm.io/gorm"
)

type Repository interface {
	DB(uow uow.UnitOfWork) *gorm.DB
}

type repositoryImpl struct{}

func (r *repositoryImpl) DB(u uow.UnitOfWork) *gorm.DB {
	log.Println("Getting the database connection from the Unit of Work")

	gormUnitOfWork, ok := u.(*UnitOfWork)
	if !ok {
		log.Fatalf("Could not convert the Unit of Work to a GORM Unit of Work")
	}

	return gormUnitOfWork.Tx
}

func NewRepository() Repository {
	return &repositoryImpl{}
}
