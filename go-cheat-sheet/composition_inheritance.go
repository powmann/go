package main

type Address struct {
	Street string
	Number int
}

type Person struct {
	Name    string
	Age     int
	Address Address // composicao de structs - Person tem um Address
}

func main() {
	p := Person{
		Name: "Jo",
		Age:  1,
		Address: Address{
			Street: "Street",
			Number: 1,
		},
	}
	println(p.Name)
	println(p.Address.Street)
}
