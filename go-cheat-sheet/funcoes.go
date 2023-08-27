package main

import (
	"fmt"
)

func main() {
	greet()
	fmt.Println(add(1, 2))
	fmt.Println(sumANDmult(2, 3))

	valum, valdois := sumANDmult(2, 3)
	fmt.Printf("sum: %v mult: %v", valum, valdois)
	fmt.Println()

	fmt.Println(swapVals("jo", "rb"))
	fmt.Println()

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
func swapVals(a, b string) (string, string) {
	return b, a
}

// retorno multiplo nomeado
func sumANDmult(a, b int) (x, y int) {
	x = a + b
	y = a * b
	return
}
