package main

import "fmt"

//Imovel é uma estrutura que define um imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := Imovel{}
	fmt.Printf("a casa é: %+v\r\n", casa)

	apartamento := Imovel{20, 40, "Meu apartamento", 200000}
	fmt.Printf("o apartamento é: %+v\r\n", apartamento)

	chacara := Imovel{
		X:     900,
		Nome:  "Minha chacara",
		valor: 800000,
		Y:     500,
	}
	fmt.Printf("a chacara é: %+v\r\n", chacara)

	fazenda := Imovel{}
	fazenda.Nome = "Minha fazenda"
	fazenda.Y = 600
	fazenda.X = 900
	fazenda.valor = 600000
	fmt.Printf("a fazenda é: %+v\r\n", fazenda)
}
