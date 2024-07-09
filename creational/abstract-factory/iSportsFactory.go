package main

import "fmt"

// Abstract factory interface

type iSportsFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}

	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
