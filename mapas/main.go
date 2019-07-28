package main

import (
	"fmt"

	"github.com/caiopapai/curso_de_go/mapas/modelo"
)

var cache map[string]modelo.Pessoa

func main() {

	cache = make(map[string]modelo.Pessoa, 0)

	funcionario1 := modelo.Pessoa{}
	funcionario1.Nome = "Caio Papai"
	funcionario2 := modelo.Pessoa{}
	funcionario2.Nome = "Jenifer Sales"

	cache[funcionario1.Nome] = funcionario1
	cache[funcionario2.Nome] = funcionario2

	fmt.Println(len(cache), cache)

	for chave, pessoa := range cache {
		if pessoa.Nome == "Caio Papai" {
			fmt.Print("O caio está nessa lista" + chave)
		}
	}

	pessoa1, existe := cache["Caio Papai"]
	if existe {
		fmt.Print("Sim, o caio está na lista" + pessoa1.Nome)
	}

	delete(cache, "Jenifer Sales")
	fmt.Print(len(cache))
}
