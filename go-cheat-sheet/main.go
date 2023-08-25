package main

import "fmt"

var age int
// var intgrande int64

var x, y int = 2, 4


var f1 float32
var f2 float64
var isTreue bool = false

tipos compostos

array
var arr [5]int
arr :=[3]int{1,2,4}

slice - arrays dinamicos

slc := []int{1,2,3,4,5,6}
subSlc := slc[1:4] // slc[1], slc[2], slc[3]

slc = append(slc, 7, 8)

maps

m := make(map[string]int)




func main() {

	fmt.Println("yeyyy")

	fmt.Println(age)
	age=11
	fmt.Println(age)

	name := "jorb"
	a, b := "jo", "rb"
	
	fmt.Println(name)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(a)
	fmt.Printf("%s%s", a, b)
	fmt.Println()
}
