package comment

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

func (s *Service) CreateComment(ctx context.Context, postID, userID uuid.UUID, content string) (*Comment, error) {
	comment := &Comment{
		ID:      uuid.New(),
		PostID:  postID,
		UserID:  userID,
		Content: content,
	}
	if err := s.Repo.Create(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *Service) ListCommentsByPost(ctx context.Context, postID uuid.UUID) ([]Comment, error) {
	return s.Repo.FindByPostID(postID)
}

func (s *Service) DeleteComment(ctx context.Context, id uuid.UUID) error {
	return s.Repo.Delete(id)
}
