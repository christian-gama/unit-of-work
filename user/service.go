package user

import (
	"errors"

	"github.com/christian-gama/uow/uow"
)

// Service is an interface that defines the functions necessary for a user service.
type Service interface {
	FindAll() ([]*User, error)
	Save(user *User) error
	TransferMoney(from uint, to uint, amount int) error
	FindOne(id uint) (*User, error)
}

type serviceImpl struct {
	repository Repository
	uow        uow.UnitOfWork
}

// NewService is a function that creates a new user service instance using the provided user repository and unit of work.
func NewService(repository Repository, uow uow.UnitOfWork) Service {
	return &serviceImpl{repository, uow}
}

// FindAll is a function that returns all the users in the database.
// It uses the user repository instance to call the appropriate function and returns the result.
func (s *serviceImpl) FindAll() ([]*User, error) {
	return s.repository.FindAll(s.uow)
}

// Save is a function that creates a new user in the database.
// It uses the user repository instance to call the appropriate function and saves the user to the database.
func (s *serviceImpl) Save(user *User) error {
	return s.repository.Save(s.uow, user)
}

// TransferMoney is a function that transfers money from one user to another.
// It uses the user repository instance to call the appropriate function and transfers money from one user to another.
func (s *serviceImpl) TransferMoney(from uint, to uint, amount int) error {
	s.uow.Begin()

	if err := s.repository.TransferMoney(s.uow, from, to, amount); err != nil {
		s.uow.Rollback()
		return err
	}

	fromUser, err := s.repository.FindOne(s.uow, from)
	if err != nil {
		s.uow.Rollback()
		return err
	}

	toUser, err := s.repository.FindOne(s.uow, to)
	if err != nil {
		s.uow.Rollback()
		return err
	}

	if fromUser.Money < 0 || toUser.Money < 0 {
		s.uow.Rollback()
		return errors.New("Cannot have a negative balance")
	}

	return s.uow.Commit()

}

// FindOne is a function that returns a user by its ID.
// It uses the user repository instance to call the appropriate function and returns the user.
func (s *serviceImpl) FindOne(id uint) (*User, error) {
	return s.repository.FindOne(s.uow, id)
}
