package user

import "github.com/christian-gama/uow/uow"

// Service is an interface that defines the functions necessary for a user service.
type Service interface {
	FindAll() ([]*User, error)
	Save(user *User) error
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
