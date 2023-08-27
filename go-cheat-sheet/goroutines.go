package main

import (
	"fmt"
	"time"
)

func main() {
	go greet()
	time.Sleep(1 * time.Second) // atraves do sleep, a goroutine tem tempo de executar
}

// goroutines - funcoes que podem ser executadas em paralelo
// goroutines sao executadas no mesmo endereco de memoria, logo, nao ha necessidade de mutex
func greet() {
	fmt.Println("gruessech!")
}
