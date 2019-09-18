package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/caiopapai/curso_de_go/get_service/model"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "Caio Augusto Papai"

	payload, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println("Erro ao converter struct para JSON", err.Error())
		return
	}

	//request, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	request, err := http.NewRequest("POST", "https://en28rz3ybbjwq.x.pipedream.net", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Erro ao realizar request", err.Error())
		return
	}

	request.Header.Set("content-type", "application/json")
	request.Header.Set("accept-language", "pt-BR,pt;q=0.8,en-US;q=0.5,en;q=0.3")
	request.SetBasicAuth("token", "token")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Erro ao ler conteudo", err.Error())
		return
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Erro ao ler conteudo", err.Error())
			return
		}
		responseBody := string(corpo)
		fmt.Println(responseBody)

	}

}
