package user

import (
	"github.com/christian-gama/uow/sql"
	"github.com/christian-gama/uow/uow"
)

// Repository is an interface that defines the functions necessary for a user repository.
type Repository interface {
	FindAll(uow uow.UnitOfWork) ([]*User, error)
	Save(uow uow.UnitOfWork, user *User) error
}

type userRepositoryImpl struct {
	sql.Repository
}

func NewRepository(repository sql.Repository) Repository {
	return &userRepositoryImpl{repository}
}

// FindAll is a function that returns all the users in the database.
// It uses the SQL repository instance to get a GORM database connection from the provided unit of work and executes a
// SELECT statement on the User table. It returns the result and any errors encountered during the query.
func (repo *userRepositoryImpl) FindAll(uow uow.UnitOfWork) ([]*User, error) {
	db := repo.DB(uow)
	var users []*User

	err := db.Find(&users).Error

	return users, err
}

// Save is a function that creates a new user in the database.
// It uses the SQL repository instance to get a GORM database connection from the provided unit of work and saves the user to the database
// using the GORM's Save function. It returns any errors encountered during the operation.
func (repo *userRepositoryImpl) Save(uow uow.UnitOfWork, user *User) error {
	db := repo.DB(uow)

	return db.Save(&user).Error
}
