package main

import (
	"encoding/json"
	"fmt"

	"github.com/caiopapai/structs_avancados/model"
)

func main() {

	casa := model.Imovel{}
	casa.X = 200
	casa.Y = 300
	casa.SetValor(3000)
	casa.Nome = "Casa do caio"

	fmt.Printf("O valor da casa é: %+v\r\n", casa.GetValor())

	objJSON, _ := json.Marshal(casa)
	fmt.Printf(string(objJSON))

}
