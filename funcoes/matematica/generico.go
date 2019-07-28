package matematica

//Calculo que executa qualquer tipo de funcao
func Calculo(funcao func(int, int) int, numeroa int, numerob int) int {
	return funcao(numeroa, numerob)
}
