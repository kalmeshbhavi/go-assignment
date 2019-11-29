package engine

import (
	"fmt"

	"github.com/kalmeshbhavi/go-assignment/domain"
	"github.com/kalmeshbhavi/go-assignment/errors"
)

var knights []domain.Knight

func (engine *arenaEngine) GetKnight(ID string) (*domain.Knight, error) {
	const op errors.Op = "engine.GetKnight"
	fighter, err := engine.knightRepository.Find(ID)
	if err != nil {
		return nil, errors.E(op, err, fmt.Sprintf("Knight #%s not found.", ID))
	}

	return fighter, nil
}

func (engine *arenaEngine) ListKnights() ([]*domain.Knight, error) {
	const op errors.Op = "engine.ListKnights"
	knights, err := engine.knightRepository.FindAll()
	if err != nil {
		return nil, errors.E(op, err, "No records found.")
	}
	return knights, nil
}

func (engine *arenaEngine) CreateKnight(knight *domain.Knight) (int64, error) {
	const op errors.Op = "engine.CreateKnight"

	id, err := engine.knightRepository.Save(knight)
	if err != nil {
		return 0, errors.E(op, err, "Unable to create knight")
	}
	return id, nil
}
