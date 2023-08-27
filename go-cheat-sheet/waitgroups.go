package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	// waitgroups - permite que uma goroutine aguarde o fim da execucao de varias outras goroutines
	// https://gobyexample.com/waitgroups
	var wg sync.WaitGroup
	wg.Add(3)

	go worker(&wg, "1")
	go worker(&wg, "2")
	go worker(&wg, "3")

	fmt.Println("waiting for workers")

	wg.Wait()
}

func worker(wg *sync.WaitGroup, nome string) {
	// wait a random amount of time between 0 and 5 seconds
	secs := rand.Intn(5)
	time.Sleep(time.Duration(secs) * time.Second)
	defer wg.Done()
	fmt.Println("worker " + nome + " done in " + strconv.Itoa(secs) + " seconds")
}
