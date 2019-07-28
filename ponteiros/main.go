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
	casa := new(Imovel)
	fmt.Printf("Casa é: %p - %+v\r\n", &casa, casa)

	apartamento := Imovel{20, 40, "Meu apartamento", 200000}
	fmt.Printf("o apartamento é:  %p - %+v\r\n", &apartamento, apartamento)

	mudaImovel(&apartamento)
	fmt.Printf("o apartamento é:  %p - %+v\r\n", &apartamento, apartamento)
}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 2
	imovel.Y = imovel.Y - 2
}
