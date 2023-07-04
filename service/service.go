package service

import "slash/models"

type UserService interface {
	CreateUser(models.User) error
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateUser(u models.User) error {
	return nil
}
