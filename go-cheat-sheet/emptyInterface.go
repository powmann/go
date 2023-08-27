package main

// empty interface
type Any interface{}

func main() {
	// empty interface
	var i Any = 1
	i = "string"
	i = true
}

// empty interface é uma interface que nao possui nenhum metodo,
// logo, qualquer tipo em Go implementa uma empty interface
