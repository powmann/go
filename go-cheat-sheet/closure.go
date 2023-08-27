package main

import (
	"fmt"
)

func main() {
	testint1 := cont()
	fmt.Println(testint1())
	fmt.Println(testint1())

	testint2 := cont()
	fmt.Println(testint2())
}

// closure - funcao que retorna uma funcao
func cont() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
