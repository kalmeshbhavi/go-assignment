package domain

type Arena struct {
}

func (arena *Arena) Fight(fighter1 Fighter, fighter2 Fighter) Fighter {

	if fighter1.GetPower() == fighter2.GetPower() {
		return nil
	}

	if fighter1.GetPower() < fighter2.GetPower() {
		return fighter2
	}

	return fighter1
}
