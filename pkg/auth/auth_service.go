package auth

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	u "github.com/waelhamzee/devconnector/pkg/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *AuthRepository
}

func NewAuthService(repo *AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(name, email, password string) (*u.User, error) {
	email = strings.ToLower(email)
	if existing, _ := s.repo.FindByEmail(email); existing != nil {
		return nil, errors.New("email already registered")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &u.User{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}

func (s *AuthService) Authenticate(email, password string) (*u.User, error) {
	email = strings.ToLower(email)
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	user.Password = ""
	return user, nil
}
