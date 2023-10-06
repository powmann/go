package main

import "fmt"

func divide(a, b int) int {
	if b == 0 {
		panic("Divide by zero")
	}
	return a / b
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fmt.Println(divide(4, 2))
	fmt.Println(divide(4, 0))
	fmt.Println("This line will not be printed")

}
