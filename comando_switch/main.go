package main

import (
	"fmt"
)

var descricao string

func main() {

	numero := 3

	switch numero {
	case 1:
		descricao = "um"
	case 2:
		descricao = "dois"
	case 3:
		descricao = "três"
	}

	fmt.Print("O numero ", numero, " se escreve assim: ", descricao)

}
