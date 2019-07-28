package main

import "fmt"

var (
		//Endereço é um valor importante e deve ser publico
		//Quando uma variável tem a primeira letra maiuscula significa que ele é publico
		Endereco, telefone string
		//O nome deve ser publico
		Nome string
 		contem bool
 		quantidade int
 		caractere rune
		 preco float64
	)


func main () {

	fmt.Println("Endereco: ", Endereco)
	fmt.Println("Contem: ", contem)
	fmt.Println("Quantidade", quantidade)
}