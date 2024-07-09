package main

import "fmt"

// Abstract Factory Design Pattern is a creational design pattern that lets you create a family of related objects. It is an abstraction over the factory pattern. It is best explained with an example. Let’s say we have two factories

// . nike
// . adidas
// Imagine you need to buy a sports kit which has a shoe and short.
// Preferably most of the time you would want to buy a full sports kit of a similar factory i.e either nike or adidas.
// This is where the abstract factory comes into the picture as concrete products
// that you want is shoe and a short and these products will be created by the abstract factory of nike and adidas.

// Both these two factories – nike and adidas implement iSportsFactory interface.
// We have two product interfaces.

// iShoe – this interface is implemented by nikeShoe and adidasShoe concrete product.
// iShort – this interface is implemented by nikeShort and adidasShort concrete product.

// iSportsFactory
// type iSportsFactory interface {
// 	makeShoe() iShoe
// 	makeShort() iShort
// }

// func getSportsFactory(brand string) (iSportsFactory, error) {
// 	if brand == "adidas" {
// 		return &adidas{}, nil
// 	}

// 	if brand == "nike" {
// 		return &nike{}, nil
// 	}

// 	return nil, fmt.Errorf("wrong brand type passed")
// }

// iShoe
// type iShoe interface {
// 	setLogo(logo string)
// 	getLogo() string
// 	setSize(size int)
// 	getSize() int
// }

// type shoe struct {
// 	logo string
// 	size int
// }

// func (s *shoe) setLogo(logo string) {
// 	s.logo = logo
// }

// func (s *shoe) getLogo() string {
// 	return s.logo
// }

// func (s *shoe) setSize(size int) {
// 	s.size = size
// }

// func (s *shoe) getSize() int {
// 	return s.size
// }

// // adidas shoe
// type adidasShoe struct {
// 	shoe
// }

// // nikeShoe
// type nikeShoe struct {
// 	shoe
// }

// iShort

// type iShort interface {
// 	setLogo(logo string)
// 	getLogo() string
// 	setSize(size int)
// 	getSize() int
// }

// type short struct {
// 	logo string
// 	size int
// }

// func (s *short) setLogo(logo string) {
// 	s.logo = logo
// }

// func (s *short) getLogo() string {
// 	return s.logo
// }

// func (s *short) setSize(size int) {
// 	s.size = size
// }

// func (s *short) getSize() int {
// 	return s.size
// }

// // adidasShort
// type adidasShort struct {
// 	short
// }

// // nikeShort
// type nikeShort struct {
// 	short
// }

// adidas
// type adidas struct{}

// func (a *adidas) makeShoe() iShoe {
// 	return &adidasShoe{
// 		shoe: shoe{
// 			logo: "adidas",
// 			size: 14,
// 		},
// 	}
// }

// func (a *adidas) makeShort() iShort {
// 	return &adidasShort{
// 		short: short{
// 			logo: "adidas",
// 			size: 14,
// 		},
// 	}
// }

// nike
// type nike struct{}

// func (n *nike) makeShoe() iShoe {
// 	return &nikeShoe{
// 		shoe: shoe{
// 			logo: "nike",
// 			size: 14,
// 		},
// 	}
// }

// func (n *nike) makeShort() iShort {
// 	return &nikeShort{
// 		short: short{
// 			logo: "nike",
// 			size: 14,
// 		},
// 	}
// }

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	adidasShoe := adidasFactory.makeShoe()
	adidasShort := adidasFactory.makeShort()

	nikeShoe := nikeFactory.makeShoe()
	nikeShort := nikeFactory.makeShort()

	printShoeDetails(adidasShoe)
	printShortDetails(adidasShort)

	printShoeDetails(nikeShoe)
	printShortDetails(nikeShort)
}

func printShoeDetails(s iShoe) {
	fmt.Println("Shoe logo:", s.getLogo())
	fmt.Println("Shoe size:", s.getSize())
}

func printShortDetails(s iShort) {
	fmt.Println("Short logo:", s.getLogo())
	fmt.Println("Short size:", s.getSize())
}
