package engine

import (
	"errors"
	"fmt"

	"github.com/kalmeshbhavi/go-assignment/domain"
)

var knights []domain.Knight

func (engine *arenaEngine) GetKnight(ID string) (*domain.Knight, error) {
	fighter := engine.knightRepository.Find(ID)
	if fighter == nil {
		return nil, errors.New(fmt.Sprintf("Knight #%s not found.", ID))
	}

	return fighter, nil
}

func (engine *arenaEngine) ListKnights() ([]*domain.Knight, error) {
	return engine.knightRepository.FindAll()
}

func (engine *arenaEngine) CreateKnight(knight *domain.Knight) (int64, error) {
	id, err := engine.knightRepository.Save(knight)
	if err != nil {
		return 0, nil
	}
	return id, nil
}
