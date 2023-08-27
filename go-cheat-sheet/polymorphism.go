package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	breed string
}

func (d Dog) Speak() string {
	return "Au au"
}

type Cat struct {
	name string
}

func (c Cat) Speak() string {
	return "Miau"
}

func main() {

	animals := []Animal{Dog{}, Cat{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	garfield := Cat{name: "Garfield"}
	fmt.Println(garfield.Speak() + " " + garfield.name)

	pluto := Dog{}
	pluto.breed = "Beagle"
	fmt.Println(pluto.Speak() + " " + pluto.breed)

}
