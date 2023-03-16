package user

import "github.com/christian-gama/uow/uow"

type Service interface {
	FindAll() ([]*User, error)
	Save(user *User) error
}

type serviceImpl struct {
	repository Repository
	uow        uow.UnitOfWork
}

func (s *serviceImpl) FindAll() ([]*User, error) {
	return s.repository.FindAll(s.uow)
}

func (s *serviceImpl) Save(user *User) error {
	return s.repository.Save(s.uow, user)
}

func NewService(repository Repository, uow uow.UnitOfWork) Service {
	return &serviceImpl{repository, uow}
}
