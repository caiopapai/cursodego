package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("[Main] Houve um erro ao abrir o arquivo: ", err.Error())
		return
	}

	scanner := bufio.NewScanner(arquivo)

	for scanner.Scan() {
		linha := scanner.Text()
		fmt.Println("O conteudo da linha é: ", linha)
	}

	arquivo.Close()

	arquivoCsv, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("[Main] Houve um erro ao abrir o arquivo: ", err.Error())
		return
	}

	leitorCsv := csv.NewReader(arquivoCsv)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[Main] Houve um erro ao abrir o arquivo: ", err.Error())
		return
	}

	for indiceLinha, linha := range conteudo {
		fmt.Printf("Linha[%d] é %s\r\n", indiceLinha, linha)
		for indiceItem, item := range linha {
			fmt.Printf("Item[%d] é %s\r\n", indiceItem, item)
		}
	}

	arquivoCsv.Close()

}
