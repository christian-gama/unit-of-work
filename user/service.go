package user

import (
	"context"
	"errors"

	"github.com/christian-gama/uow/uow"
)

// Service is an interface that defines the functions necessary for a user service.
type Service interface {
	FindAll(ctx context.Context, params *FindAllParams) ([]*User, error)
	Save(ctx context.Context, params *SaveParams) error
	TransferMoney(ctx context.Context, params *TransferMoneyParams) error
	FindOne(ctx context.Context, params *FindOneParams) (*User, error)
	Delete(ctx context.Context, params *DeleteParams) error
}

type serviceImpl struct {
	repo Repository
	uow  uow.UnitOfWork
}

// NewService is a function that creates a new user service instance using the provided user repository and unit of work.
func NewService(repository Repository, uow uow.UnitOfWork) Service {
	return &serviceImpl{repository, uow}
}

// FindAll is a function that returns all the users in the database.
// It uses the user repository instance to call the appropriate function and returns the result.
func (s *serviceImpl) FindAll(ctx context.Context, params *FindAllParams) ([]*User, error) {
	return s.repo.FindAll(ctx, params, nil)
}

// Save is a function that creates a new user in the database.
// It uses the user repository instance to call the appropriate function and saves the user to the database.
func (s *serviceImpl) Save(ctx context.Context, params *SaveParams) error {
	return s.repo.Save(ctx, params, nil)
}

// TransferMoney is a function that transfers money from one user to another.
// It uses the user repository instance to call the appropriate function and transfers money from one user to another.
func (s *serviceImpl) TransferMoney(ctx context.Context, params *TransferMoneyParams) error {
	return s.uow.Transaction(func(uow uow.UnitOfWork) error {
		if err := s.repo.TransferMoney(ctx, params, uow); err != nil {
			return err
		}

		fromUser, err := s.repo.FindOne(ctx, &FindOneParams{ID: params.From}, uow)
		if err != nil {
			return err
		}

		toUser, err := s.repo.FindOne(ctx, &FindOneParams{ID: params.To}, uow)
		if err != nil {
			return err
		}

		if fromUser.Money < 0 || toUser.Money < 0 {
			return errors.New("cannot have a negative balance")
		}

		return nil
	})
}

// FindOne is a function that returns a user by its ID.
// It uses the user repository instance to call the appropriate function and returns the user.
func (s *serviceImpl) FindOne(ctx context.Context, params *FindOneParams) (*User, error) {
	return s.repo.FindOne(ctx, params, nil)
}

// Delete is a function that deletes a user by its ID.
// It uses the user repository instance to call the appropriate function and deletes the user.
func (s *serviceImpl) Delete(ctx context.Context, params *DeleteParams) error {
	return s.repo.Delete(ctx, params, nil)
}
