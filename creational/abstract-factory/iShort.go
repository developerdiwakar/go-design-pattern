package main

type iShort interface {
	setLogo(logo string)
	getLogo() string
	setSize(size int)
	getSize() int
}

type short struct {
	logo string
	size int
}

func (s *short) setLogo(logo string) {
	s.logo = logo
}

func (s *short) getLogo() string {
	return s.logo
}

func (s *short) setSize(size int) {
	s.size = size
}

func (s *short) getSize() int {
	return s.size
}

// adidasShort
type adidasShort struct {
	short
}

// nikeShort
type nikeShort struct {
	short
}