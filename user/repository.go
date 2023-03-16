package user

import (
	"github.com/christian-gama/uow/sql"
	"github.com/christian-gama/uow/uow"
	"gorm.io/gorm"
)

// Repository is an interface that defines the functions necessary for a user repository.
type Repository interface {
	FindAll(uow uow.UnitOfWork) ([]*User, error)
	Save(uow uow.UnitOfWork, user *User) error
	TransferMoney(uow uow.UnitOfWork, from uint, to uint, amount int) error
	FindOne(uow uow.UnitOfWork, id uint) (*User, error)
}

type userRepositoryImpl struct {
	sql.Repository
}

// NewRepository is a function that returns a new user repository instance.
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

// TransferMoney is a function that transfers money from one user to another.
// It uses the SQL repository instance to get a GORM database connection from the provided unit of work and executes a
// UPDATE statement on the User table. It returns any errors encountered during the operation.
func (repo *userRepositoryImpl) TransferMoney(uow uow.UnitOfWork, from uint, to uint, amount int) error {
	db := repo.DB(uow)

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&User{}).Where("id = ?", from).Update("money", gorm.Expr("money - ?", amount)).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("id = ?", to).Update("money", gorm.Expr("money + ?", amount)).Error; err != nil {
			return err
		}

		return nil
	})
}

// FindOne is a function that returns a user from the database by its ID.
// It uses the SQL repository instance to get a GORM database connection from the provided unit of work and executes a
// SELECT statement on the User table. It returns the result and any errors encountered during the query.
func (repo *userRepositoryImpl) FindOne(uow uow.UnitOfWork, id uint) (*User, error) {
	db := repo.DB(uow)

	var user User

	err := db.First(&user, id).Error

	return &user, err
}
