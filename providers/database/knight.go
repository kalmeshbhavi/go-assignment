package database

import (
	"github.com/kalmeshbhavi/go-assignment/domain"
)

type knightRepository struct {
	provider *Provider
}

func (repository *knightRepository) FindAll() ([]*domain.Knight, error) {

	var knights []*domain.Knight
	rows, err := repository.provider.DB.Query("SELECT ID, NAME, STRENGTH, WEAPON_POWER FROM KNIGHTS")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		knight := domain.Knight{}
		err := rows.Scan(&knight.ID, &knight.Name, &knight.Strength, &knight.WeaponPower)
		if err != nil {
			return nil, err
		}
		knights = append(knights, &knight)
	}
	return knights, nil
}

func (repository *knightRepository) Save(knight *domain.Knight) (int64, error) {
	stmt, err := repository.provider.DB.Prepare("INSERT INTO knights(name, strength, weapon_power) VALUES(?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(knight.Name, knight.Strength, knight.WeaponPower)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (repository *knightRepository) Find(ID string) *domain.Knight {

	rows := repository.provider.DB.QueryRow("SELECT ID, NAME, STRENGTH, WEAPON_POWER FROM KNIGHTS WHERE ID = ?", ID)

	knight := domain.Knight{}
	err := rows.Scan(&knight.ID, &knight.Name, &knight.Strength, &knight.WeaponPower)
	if err != nil {
		return nil
	}

	return &knight
}
