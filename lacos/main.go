package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println("O valor de i é: ", i)
	}

	valor := 0
	teste := true

	for teste {
		if valor == 0 {
			fmt.Println("Entrei no laço")
			teste = false
		}
	}

}
