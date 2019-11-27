package domain

type Fighter interface {
	GetID() string
	GetPower() float64
}

type Knight struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Strength    int    `json:"strength"`
	WeaponPower int    `json:"weapon_power"`
}
