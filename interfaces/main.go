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

// SomEmitidoPorGalinha é um som
func SomEmitidoPorGalinha(g model.Ave) {
	fmt.Println(g.Cacareja())
}

// SomEmitidoPorPato é um som
func SomEmitidoPorPato(p model.Ave) {
	fmt.Println(p.Quack())
}
