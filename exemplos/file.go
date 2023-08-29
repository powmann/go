package main

import (
	"fmt"
	"os"
)

func main() {

	// escreve no arquivo
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("file.txt", d1, 0644)
	check(err)

	// le o arquivo
	dat, err := os.ReadFile("file.txt")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("file.txt")
	check(err)

	f.Close()

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
