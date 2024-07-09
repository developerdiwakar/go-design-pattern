package main

// Concrete factory

type adidas struct{}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *adidas) makeShort() iShort {
	return &adidasShort{
		short: short{
			logo: "adidas",
			size: 34,
		},
	}
}
