package engine

import (
	"github.com/kalmeshbhavi/go-assignment/domain"
	"github.com/kalmeshbhavi/go-assignment/providers/database"
)

type Engine interface {
	GetKnight(ID string) (*domain.Knight, error)
	ListKnights() ([]*domain.Knight, error)
	CreateKnight(*domain.Knight) (int64, error)
}

type arenaEngine struct {
	arena            *domain.Arena
	knightRepository database.KnightRepository
}

func NewEngine(db database.DatabaseProvider) Engine {
	return &arenaEngine{
		arena:            &domain.Arena{},
		knightRepository: db.GetKnightRepository(),
	}
}
