package user

import (
	"log"

	"github.com/christian-gama/uow/sql"
	"github.com/christian-gama/uow/uow"
)

type Repository interface {
	FindAll(uow uow.UnitOfWork) ([]*User, error)
	Save(uow uow.UnitOfWork, user *User) error
}

type userRepositoryImpl struct {
	sql.Repository
}

func (repo *userRepositoryImpl) FindAll(uow uow.UnitOfWork) ([]*User, error) {
	db := repo.DB(uow)
	var users []*User

	log.Printf("Searching for all users")

	err := db.Find(&users).Error

	return users, err
}

func (repo *userRepositoryImpl) Save(uow uow.UnitOfWork, user *User) error {
	db := repo.DB(uow)

	log.Printf("Saving user %s to the Database", user.Name)

	return db.Save(&user).Error
}

func NewRepository(repository sql.Repository) Repository {
	return &userRepositoryImpl{repository}
}
