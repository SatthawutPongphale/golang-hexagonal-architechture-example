package application

import (
	"golang-hexagonal-architechture-example/internal/domain"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(input domain.UserInput) (*domain.User, error) {
	user := &domain.User{
		Name:  input.Name,
		Email: input.Email,
	}
	return s.repo.Save(user)
}
