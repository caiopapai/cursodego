package main

import "fmt"

func main() {
	var nomes []string
	nomes = make([]string, 10)
	nomes[0] = "Caio Augusto Papai"

	// Inclui e aumenta o slice
	nomes[1] = "João da Silda"
	nomes[2] = "Maria da Silda"
	nomes[3] = "Paulo da Silda"
	nomes = append(nomes, "Marco marcola")
	// Tamanho do Slice
	fmt.Println(len(nomes))
	// Capacidade do Slice
	fmt.Println(cap(nomes))
	// Pega um intervalo
	listaDeNomes := nomes[0:5]
	fmt.Println(listaDeNomes)
	// Pega pelo começo
	temp1 := nomes[:2]
	fmt.Println(temp1)
	// Pega pelo final
	temp2 := nomes[len(nomes)-1:]
	fmt.Println(temp2)
	// Removendo itens do slice
	itensARetirar := 2
	temp := nomes[:itensARetirar]
	temp = append(temp, nomes[itensARetirar+1:]...)
	copy(nomes, temp)
	fmt.Println(nomes)

}
