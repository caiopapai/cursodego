package main

import (
	"fmt"

	"github.com/caiopapai/curso_de_go/interfaces/model"
)

func main() {

	aveRara := model.Ave{}
	aveRara.Nome = "Ave rara especial"

	galinha := model.Ave{}
	galinha.Nome = "Galinha"

	pato := model.Ave{}
	pato.Nome = "Pato"

	SomEmitidoPorGalinha(galinha)
	SomEmitidoPorPato(pato)

	SomEmitidoPorPato(aveRara)
	SomEmitidoPorGalinha(aveRara)
}

func SomEmitidoPorGalinha(g model.Ave) {
	fmt.Println(g.Cacareja())
}

func SomEmitidoPorPato(p model.Ave) {
	fmt.Println(p.Quack())
}
