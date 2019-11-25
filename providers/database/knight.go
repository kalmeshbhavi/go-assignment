package database

import "github.com/kalmeshbhavi/go-assignment/domain"

type knightRepository struct{}

func (repository *knightRepository) FindAll() []*domain.Knight {
	panic("implement me")
}

func (repository *knightRepository) Save(knight *domain.Knight) {
	panic("implement me")
}

func (repository *knightRepository) Find(ID string) *domain.Knight {
	return nil
}
