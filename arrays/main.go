package main

import "fmt"

func main() {

	var teste [3]int
	teste[0] = 1
	teste[1] = 2
	teste[2] = 3

	nomes := [2]string{"Caio Augusto Papai", "Jenifer Rodrigues Sales"}
	estados := [...]string{"São Paulo", "Rio de Janeiro", "Minas Gerais"}

	fmt.Println(nomes)

	for indice, estado := range estados {
		fmt.Printf("Estado[%d] é %s\n\r", indice, estado)
	}

}
