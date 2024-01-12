package tournament

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

type (
	Service interface {
		Create(ctx context.Context, title string) (*Tournament, error)

		Get(ctx context.Context, id uuid.UUID) (*Tournament, error)

		Join(ctx context.Context, request JoinRequest) (*Tournament, error)
	}

	MockService struct {
		tournaments []*Tournament
	}
)

func NewService() Service {
	return &MockService{tournaments: make([]*Tournament, 0)}
}

func (service *MockService) Create(ctx context.Context, title string) (*Tournament, error) {
	t := New(title)
	service.tournaments = append(service.tournaments, t)
	return t, nil
}

func (service *MockService) Get(ctx context.Context, id uuid.UUID) (*Tournament, error) {
	for _, t := range service.tournaments {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, errors.New("tournament not found")
}

func (service *MockService) Join(ctx context.Context, request JoinRequest) (*Tournament, error) {
	return nil, nil
}
