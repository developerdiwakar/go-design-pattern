package main

import "fmt"

//Introduction:
// Factory design pattern is a creational design pattern and it is also one of the most commonly used pattern.
// This pattern provides a way to hide the creation logic of the instances being created.
// The client only interacts with a factory struct and tells the kind of instances that needs to be created.
// The factory class interacts with the corresponding concrete structs and returns the correct instance back.
// In below example

// We have iGun interface which defines all methods a gun should have
// There is gun struct that implements the iGun interface.
// Two concrete guns ak47 and maverick. Both embed gun struct and hence also indirectly implement all methods of iGun and hence are of iGun type
// We have a gunFactory struct which creates the gun of type ak47 or maverick.
// The main.go acts as a client and instead of directly interacting with ak47 or maverick, it relies on gunFactory to create instances of ak47 and maverick

// iGun.go
// type iGun interface {
// 	setName(name string)
// 	getName() string
// 	setPower(power int)
// 	getPower() int
// }

// gun.go
// type gun struct {
// 	name  string
// 	power int
// }

// func (g *gun) SetName(name string) {
// 	g.name = name
// }

// func (g *gun) GetName() string {
// 	return g.name
// }

// func (g *gun) SetPower(power int) {
// 	g.power = power
// }

// func (g *gun) GetPower() int {
// 	return g.power
// }

// ak47.go
// type ak47 struct {
// 	gun
// }

// func newAk47() IGun {
// 	return &ak47{
// 		gun: gun{
// 			name:  "AK47 Terminus",
// 			power: 4,
// 		},
// 	}
// }

// maverick.go
// type maverick struct {
// 	gun
// }

// func newMaverick() IGun {
// 	return &maverick{
// 		gun: gun{
// 			name:  "Maverick Terminator 007",
// 			power: 5,
// 		},
// 	}
// }

// gunFactory.go
// func getGun(gunType string) (iGun, error) {
// 	if gunType == "ak47" {
// 		return newAk47(), nil
// 	}

// 	if gunType == "maverick" {
// 		return newMaverick(), nil
// 	}

// 	return nil, fmt.Errorf("wrong gun type passed")
// }

func main() {
	ak47, _ := GetGun("ak47")
	maverick, _ := GetGun("maverick")

	printDetails(ak47)
	printDetails(maverick)

}

func printDetails(g IGun) {
	fmt.Println("Gun name:", g.GetName())
	fmt.Println("Gun Power:", g.GetPower())
}
