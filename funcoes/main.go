package main

import (
	"fmt"

	"github.com/caiopapai/funcoes/matematica"
)

func main() {
	resultado := matematica.Multiplicador(2, 2)
	fmt.Printf("O resultado da multiplicacao é: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma é: %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Multiplicador, 5, 5)
	fmt.Printf("O resultado da multiplicacao com função é: %d\r\n", resultado)
	resultado = matematica.Divisor(10, 3)
	fmt.Printf("O resultado da divisão é: %d\r\n", resultado)
	resultado, resto := matematica.DivisorComResto(20, 5)
	fmt.Printf("O resultado da divisão é: %d e o resto da divisão foi: %d\n", resultado, resto)
}
