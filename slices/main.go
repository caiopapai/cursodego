package main

import "fmt"

func main() {
	var nomes []string
	nomes = make([]string, 10)
	nomes[0] = "Caio Augusto Papai"
	nomes = append(nomes, "Outro nome qualquer")
	// Tamanho do Slice
	fmt.Print(len(nomes))

	// Capacidade do Slice
	fmt.Print(cap(nomes))
}
