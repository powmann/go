package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("escrevendo do qrquivo main.go")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("jorb@jorb.online")

	fmt.Println(erro)
}
