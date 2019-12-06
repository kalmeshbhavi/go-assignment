package database

import "github.com/kalmeshbhavi/go-assignment/domain"

type KnightRepository interface {
	Find(ID string) (*domain.Knight, error)
	FindAll() ([]*domain.Knight, error)
	Save(knight *domain.Knight) (int64, error)
}
