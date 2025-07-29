package post

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreatePost(ctx context.Context, userID uuid.UUID, content string) (*Post, error) {
	post := &Post{
		ID:      uuid.New(),
		UserID:  userID,
		Content: content,
	}
	if err := s.Repo.Create(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *Service) GetPostByID(ctx context.Context, id uuid.UUID) (*Post, error) {
	return s.Repo.FindByID(id)
}

func (s *Service) ListPosts(ctx context.Context) ([]Post, error) {
	return s.Repo.FindAll()
}

func (s *Service) DeletePost(ctx context.Context, id uuid.UUID) error {
	return s.Repo.Delete(id)
}
