package main

import "fmt"
import "github.com/caiopapai/variaveis/pacotes/prefixo"

//Nome do usuário do pacote main
var NomeDoUsuario = "Caio Augusto Papai"

func main ()  {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Capital: %d\r\n", prefixo.Capital)
}