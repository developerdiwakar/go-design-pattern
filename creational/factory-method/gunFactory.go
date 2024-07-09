package main

import "fmt"

type IGun interface {
	SetName(name string)
	GetName() string
	SetPower(power int)
	GetPower() int
}

func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAk47(), nil
	}

	if gunType == "maverick" {
		return NewMaverick(), nil
	}

	return nil, fmt.Errorf("wrong gun type passed")
}
