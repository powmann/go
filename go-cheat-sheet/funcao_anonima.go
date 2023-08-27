package main

import (
	"fmt"
)

func main() {
	// funcao anonima - criancao de funcao anonima e atribuicao a uma variavel
	soma := func(a, b int) int {
		return a + b
	}

	result := soma(2, 3)
	fmt.Println(result)
}
