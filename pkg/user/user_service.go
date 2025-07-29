package user

import (
	"github.com/google/uuid"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id uuid.UUID) (*User, error) {
	return s.repo.FindUserByID(id)
}
