package connection

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

func (s *Service) CreateConnection(ctx context.Context, userID, targetID uuid.UUID) error {
	conn := &Connection{
		ID:       uuid.New(),
		UserID:   userID,
		TargetID: targetID,
	}
	return s.Repo.Create(conn)
}

func (s *Service) ListConnections(ctx context.Context, userID uuid.UUID) ([]Connection, error) {
	return s.Repo.FindByUserID(userID)
}

func (s *Service) DeleteConnection(ctx context.Context, userID, targetID uuid.UUID) error {
	return s.Repo.Delete(userID, targetID)
}
