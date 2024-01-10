package tournament

import (
	"errors"
	"github.com/google/uuid"
)

type (
	Service interface {
		Create(title string) (*Tournament, error)
		Get(id uuid.UUID) (*Tournament, error)
		Join(request JoinRequest) (*Tournament, error)
	}

	MockService struct {
		tournaments []*Tournament
	}
)

func NewService() Service {
	return &MockService{tournaments: make([]*Tournament, 0)}
}

func (service *MockService) Create(title string) (*Tournament, error) {
	t := New(title, 0)
	service.tournaments = append(service.tournaments, t)
	return t, nil
}

func (service *MockService) Get(id uuid.UUID) (*Tournament, error) {
	for _, t := range service.tournaments {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, errors.New("tournament not found")
}

func (service *MockService) Join(request JoinRequest) (*Tournament, error) {
	for _, t := range service.tournaments {
		if t.ID == request.ID && t.Code == request.Code {

		}
	}
	return nil, nil
}
