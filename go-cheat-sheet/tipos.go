package main

import (
	"fmt"
)

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

	type Coordenadas struct {
		lat, long float64
	}

	var mc map[string]Coordenadas
	mc = make(map[string]Coordenadas)

	mc["algumLugar"] = Coordenadas{22.222, 33.333}

	fmt.Println(mc["algumLugar"])
	fmt.Println(mc["algumLugar"].long)

}
