package main

import "fmt"

func main() {
	// invocacao do metodo
	p1 := Persona{"jo", "rb"}
	fmt.Println(p1.FullName())
	fmt.Println(p1.PersonaSayHi(p1))
}

// metodos - é possivel relacionar uma estrutura a uma funcao, tornando a um metodo deste tipo

type Persona struct {
	nome      string
	sobrenome string
}

func (p Persona) FullName() string {
	return p.nome + " " + p.sobrenome
}

func (p Persona) PersonaSayHi(persona Persona) string {
	return "hi, my name is " + persona.nome + persona.sobrenome
}
