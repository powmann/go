package main

import "fmt"

func main() {

	var age int // declaracao explicita
	// var graandeInt int64

	fmt.Println(age) // valor padrao
	age = 11         // atribuicao direta
	fmt.Println(age)
	fmt.Println()

	var x, y int = 2, 4 // atribuicao multipla
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println()

	name := "jorb" // inferencia de tipo
	a, b := "jo", "rb"

	fmt.Println(name)

	fmt.Println(a)
	fmt.Printf("%s%s", a, b) // print formatado https://pkg.go.dev/fmt
	fmt.Println()

	var f1 float32 = 34534.14535453
	var f2 float64 = 45.824987242342
	fmt.Printf("float32:%4.2f float64:%e", f1, f2)
	fmt.Println()

	var isTreue bool = false
	fmt.Printf("yey/ney: %t", isTreue)
	fmt.Println()

	// tipos compostos

	// array
	var arr [5]int
	arr1 := [3]int{1, 2, 4}

	fmt.Println(arr)
	fmt.Println(arr1[1])
	fmt.Println()

	// slice - arrays dinamicos

	slc := []int{1, 2, 3, 4, 5, 6}
	subSlc := slc[1:4] // slc[1], slc[2], slc[3]

	slc = append(slc, 7, 8)

	fmt.Println(subSlc)
	fmt.Println(slc)
	fmt.Println()

	// maps

	m := make(map[string]string)

	m["key1"] = "val1"
	m["key2"] = "2"

	fmt.Println(m)
	fmt.Println(m["key2"])

	// structs

	type Person struct {
		name       string
		someintval int
	}

	person1 := Person{"jo", 1}
	person2 := Person{"rb", 2}

	fmt.Println(person1.name)
	fmt.Println(person2)

	// exemplo pratico combinanndo maps e structs

	type Coords struct {
		lat, long float64
	}

	var mc map[string]Coords
	mc = make(map[string]Coords)

	mc["keyqqr"] = Coords{22.222, 33.333}

	fmt.Println(mc["keyqqr"])
	fmt.Println(mc["keyqqr"].long)


	greet()
	fmt.Println(add(1, 2))
	fmt.Println(sumANDmult(2, 3))

	valum, valdois := sumANDmult(2, 3)
	fmt.Printf("sum: %v mult: %v", valum, valdois)
	fmt.Println()

	fmt.Println(swapVals("jo", "rb"))
	fmt.Println()

	// invocacao do metodo
	p1 := Persona{"jo", "rb"}
	fmt.Println(p1.FullName())
	fmt.Println(p1.PersonaSayHi(p1))

}

// mais besta impossivel
func greet() {
	fmt.Println("Gruessech!!")
}

// funcao com retorno simples
func add(x int, y int) int {
	return x + y
}

// retorno multiplo
func swapVals (a, b string) (string, string){
	return b, a
}

// retorno multiplo nomeado
func sumANDmult (a, b int) (x, y int) {
	x = a + b
	y = a * b
	return
}

// metodos - é possivel relacionar uma estrutura a uma funcao, tornando a um metodo deste tipo

type Persona struct {
	nome string
	sobrenome string
}

func (p Persona) FullName() string {
	return p.nome + " " + p.sobrenome
}

func (p Persona) PersonaSayHi(persona Persona) string {
	return "hi, my name is " + persona.nome
}