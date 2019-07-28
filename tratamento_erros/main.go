package main

import (
	"fmt"

	"github.com/caiopapai/curso_de_go/tratamento_erros/model"
)

func main() {
	pessoa := model.Pessoa{}
	pessoa.SetNome("Caio Papai")
	if err := pessoa.SetIdade(17); err != nil {
		print(err.Error())
	}
	pessoa.SetTelefone("971441406")

	fmt.Print(pessoa)

}
