package main

type Maverick struct {
	Gun
}

func NewMaverick() IGun {
	return &Maverick{
		Gun: Gun{
			Name:  "Maverick Terminator 007",
			Power: 5,
		},
	}
}
