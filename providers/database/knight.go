package database

import (
	"strconv"

	"github.com/kalmeshbhavi/go-assignment/domain"
	"github.com/kalmeshbhavi/go-assignment/errors"
)

type knightRepository struct {
	provider *Provider
}

func (repository *knightRepository) FindAll() ([]*domain.Knight, error) {
	const op errors.Op = "knight.FindAll"
	var knights []*domain.Knight
	rows, err := repository.provider.DB.Query("SELECT ID, NAME, STRENGTH, WEAPON_POWER FROM knights")
	if err != nil {
		return nil, errors.E(op, errors.KindNotFound, err, "Unable to fetch the knight details")
	}

	for rows.Next() {
		knight := domain.Knight{}
		err := rows.Scan(&knight.ID, &knight.Name, &knight.Strength, &knight.WeaponPower)
		if err != nil {
			return nil, errors.E(op, errors.KindUnexpected, err, "Unable to parse knight db domain")
		}
		knights = append(knights, &knight)
	}
	return knights, nil
}

func (repository *knightRepository) Save(knight *domain.Knight) (int64, error) {
	const op errors.Op = "knight.Save"
	stmt, err := repository.provider.DB.Prepare("INSERT INTO knights(name, strength, weapon_power) VALUES(?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		return 0, errors.E(op, errors.KindUnexpected, err, "Error while creating insert knight sql")
	}

	result, err := stmt.Exec(knight.Name, knight.Strength, knight.WeaponPower)
	if err != nil {
		return 0, errors.E(op, errors.KindUnexpected, err, "Error while storing knight details")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.E(op, err, errors.KindUnexpected, "Error while fetching stored ID")
	}

	return id, nil
}

func (repository *knightRepository) Find(ID string) (*domain.Knight, error) {
	const op errors.Op = "knight.Find"
	i, _ := strconv.Atoi(ID)
	rows := repository.provider.DB.QueryRow("SELECT ID, NAME, STRENGTH, WEAPON_POWER FROM knights WHERE ID = ?", i)

	knight := domain.Knight{}
	err := rows.Scan(&knight.ID, &knight.Name, &knight.Strength, &knight.WeaponPower)
	if err != nil {
		return nil, errors.E(op, errors.KindNotFound, err, "Error while finding knight.")
	}

	return &knight, nil
}
