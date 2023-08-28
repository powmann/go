package main

// https://gobyexample.com/errors
import (
	"errors"
	"fmt"
)

//forma mais simples, apenas lancando um novo erro
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("nops 42 nao")
	}
	return arg + 3, nil
}

// criando um erro customizado
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// aplicando o erro customizado
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "nops denovo 42 nao"}
	}
	return arg + 3, nil
}


func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failled", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failled", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	// _, e := f2(42)
	// if ae, ok := e.(*argError); ok {
	// 	fmt.Println(ae.arg)
	// 	fmt.Println(ae.prob)
	// }
	fmt.Println()
	resp, errmsg := f2(42)

	fmt.Println(resp)
	fmt.Println(errmsg)
}
