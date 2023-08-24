package factory

import "fmt"

func GetGun(gunType string) (IGun, error) {
	if gunType == Ak47Type {
		return NewAk47(), nil
	}
	if gunType == MusketType {
		return NewMusket(), nil
	}
	if gunType == DesertEagleType {
		return NewDeserEagle(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
