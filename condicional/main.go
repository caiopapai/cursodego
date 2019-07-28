package main

import (
	"fmt"
)

func main() {
	meses := 12
	situacao := true
	cidade := "São Paulo"

	//< > <= >= && || != ==
	if meses <= 12 {
		fmt.Printf("Esse credor deve a pouco tempo")
	}

	if situacao {
		fmt.Printf("Ele está devendo")
	}

	if !situacao {
		fmt.Printf("Ele está adiplente")
	}

	if cidade == "São Paulo" {
		fmt.Printf("Cliente vive em São Paulo")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("O cliente esta devendo", descricao)
	}
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 1 {
		status = true
		descricao = "O cliente está devendo"
		return
	}

	descricao = "O cliente está com a mensalidade em dia"
	return
}
