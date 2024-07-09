package main

// Concrete factory

type nike struct{}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 12,
		},
	}
}

func (n *nike) makeShort() iShort {
	return &nikeShort{
		short: short{
			logo: "nike",
			size: 32,
		},
	}
}
