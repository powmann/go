package main

import (
	"fmt"
	"time"
)

func main() {
	// channels - canais de comunicacao entre goroutines
	ch := make(chan string)

	go func() {
		ch <- "pong"
	}()

	msg := <-ch
	fmt.Println(msg)

	// buffered channels
	// canais com buffer sao assincronos, ou seja, nao ha necessidade de uma goroutine esperar a outra para executar
	ch1 := make(chan int, 2) // buffer de 2 inteiros
	ch1 <- 144
	ch1 <- 455
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	// select
	// permite que uma goroutine aguarde varias operacoes de comunicacao, semelhante ao switch mas para channels
	chann1 := make(chan string)
	chann2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chann1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		chann2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chann1:
			fmt.Println("received", msg1)
		case msg2 := <-chann2:
			fmt.Println("received", msg2)
		}
	}
}
