package engine

import (
	"errors"
	"fmt"
	"github.com/kalmeshbhavi/go-assignment/domain"
)

func (engine *arenaEngine) GetKnight(ID string) (*domain.Knight, error) {
	fighter := engine.knightRepository.Find(ID)
	if fighter == nil {
		return nil, errors.New(fmt.Sprintf("fighter with ID '%s' not found!", ID))
	}

	return fighter, nil
}

func (engine *arenaEngine) ListKnights() []*domain.Knight {
	return nil
}

func (engine *arenaEngine) Fight(fighter1ID string, fighter2ID string) domain.Fighter {
	return nil
}
