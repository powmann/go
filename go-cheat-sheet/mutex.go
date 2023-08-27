package main

import (
	"fmt"
	"sync"
)

/* mutex - Um mutex (mutual exclusion) é uma estrutura de sincronização usada para
controlar o acesso concorrente a recursos compartilhados, garantindo que apenas uma
goroutine (thread) tenha acesso a esses recursos por vez. */

func main() {
	var contador int
	var mutex sync.Mutex

	// Função para incrementar o contador de forma segura usando um mutex
	incrementar := func() {
		mutex.Lock()
		defer mutex.Unlock()
		contador++
	}

	// Criando um grupo de espera para as goroutines
	var wg sync.WaitGroup

	// Executando várias goroutines que incrementam o contador
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrementar()
		}()
	}

	// Aguardando a conclusão de todas as goroutines
	wg.Wait()

	fmt.Printf("Contador final: %d\n", contador)
}

// https://gobyexample.com/mutexes
