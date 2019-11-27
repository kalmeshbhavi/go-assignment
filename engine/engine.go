package engine

import "github.com/kalmeshbhavi/go-assignment/domain"

type Engine interface {
	GetKnight(ID string) (*domain.Knight, error)
	ListKnights() ([]*domain.Knight, error)
	CreateKnight(*domain.Knight) (int64, error)
}

type KnightRepository interface {
	Find(ID string) *domain.Knight
	FindAll() ([]*domain.Knight, error)
	Save(knight *domain.Knight) (int64, error)
}

type DatabaseProvider interface {
	GetKnightRepository() KnightRepository
}

type arenaEngine struct {
	arena            *domain.Arena
	knightRepository KnightRepository
}

func NewEngine(db DatabaseProvider) Engine {
	return &arenaEngine{
		arena:            &domain.Arena{},
		knightRepository: db.GetKnightRepository(),
	}
}
