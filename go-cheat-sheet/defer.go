package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func readFile(filename string) {

	file, error := os.Open(filename)

	if error != nil {
		log.Fatalf("Error opening file: %s", error)
	}

	defer file.Close()

	dat, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %s", error)
	}
	fmt.Print(string(dat))

	// Do something with the file
	content := make([]byte, 32)
	for {
		_, error := file.Read(content)

		if error != nil {
			log.Fatalf("Error reading file: %s", error)
		}

		// Do something with the content
		for _, c := range content {
			// Do something with c
			fmt.Println(c)
		}

		if error == io.EOF {
			break
		}
	}
}

func main() {

	readFile("../exemplos/file.txt")

}
