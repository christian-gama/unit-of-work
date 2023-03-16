package user

import (
	"context"

	"github.com/christian-gama/uow/sql"
	"github.com/christian-gama/uow/uow"
	"gorm.io/gorm"
)

// Repository is an interface that defines the functions necessary for a user repository.
type Repository interface {
	FindAll(ctx context.Context, params *FindAllParams, uow uow.UnitOfWork) ([]*User, error)
	Save(ctx context.Context, params *SaveParams, uow uow.UnitOfWork) error
	TransferMoney(ctx context.Context, params *TransferMoneyParams, uow uow.UnitOfWork) error
	FindOne(ctx context.Context, params *FindOneParams, uow uow.UnitOfWork) (*User, error)
	Delete(ctx context.Context, params *DeleteParams, uow uow.UnitOfWork) error
}

type userRepositoryImpl struct {
	sql.Repository
}

// NewRepository is a function that returns a new user repository instance.
func NewRepository(repository sql.Repository) Repository {
	return &userRepositoryImpl{repository}
}

type FindAllParams struct{}

// FindAll is a function that returns all the users in the database.
// It uses the SQL repository instance to get a GORM database connection from the provided unit of work and executes a
// SELECT statement on the User table. It returns the result and any errors encountered during the query.
func (repo *userRepositoryImpl) FindAll(ctx context.Context, params *FindAllParams, uow uow.UnitOfWork) ([]*User, error) {
	db := repo.DB(ctx, uow)

	var users []*User
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

type SaveParams struct {
	User *User
}

// Save is a function that creates a new user in the database.
// It uses the SQL repository instance to get a GORM database connection from the provided unit of work and saves the user to the database
// using the GORM's Save function. It returns any errors encountered during the operation.
func (repo *userRepositoryImpl) Save(ctx context.Context, params *SaveParams, uow uow.UnitOfWork) error {
	db := repo.DB(ctx, uow)

	return db.Save(&params.User).Error
}

type TransferMoneyParams struct {
	From   uint
	To     uint
	Amount int
}

// TransferMoney is a function that transfers money from one user to another.
// It uses the SQL repository instance to get a GORM database connection from the provided unit of work and executes a
// UPDATE statement on the User table. It returns any errors encountered during the operation.
func (repo *userRepositoryImpl) TransferMoney(ctx context.Context, params *TransferMoneyParams, uow uow.UnitOfWork) error {
	db := repo.DB(ctx, uow)

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&User{}).Where("id = ?", params.From).Update("money", gorm.Expr("money - ?", params.Amount)).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("id = ?", params.To).Update("money", gorm.Expr("money + ?", params.Amount)).Error; err != nil {
			return err
		}

		return nil
	})
}

type FindOneParams struct {
	ID uint
}

// FindOne is a function that returns a user from the database by its ID.
// It uses the SQL repository instance to get a GORM database connection from the provided unit of work and executes a
// SELECT statement on the User table. It returns the result and any errors encountered during the query.
func (repo *userRepositoryImpl) FindOne(ctx context.Context, params *FindOneParams, uow uow.UnitOfWork) (*User, error) {
	db := repo.DB(ctx, uow)

	var user *User
	err := db.First(&user, params.ID).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

type DeleteParams struct {
	ID uint
}

// Delete is a function that deletes a user from the database by its ID.
// It uses the SQL repository instance to get a GORM database connection from the provided unit of work and executes a
// DELETE statement on the User table. It returns any errors encountered during the operation.
func (repo *userRepositoryImpl) Delete(ctx context.Context, params *DeleteParams, uow uow.UnitOfWork) error {
	db := repo.DB(ctx, uow)

	return db.Delete(&User{}, params.ID).Error
}
