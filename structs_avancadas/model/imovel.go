package model

//Imovel é uma estrutura que define um imovel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//SetValor é o setter de valor da estrutura de imovel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//GetValor é o getter do valor da estrutura de imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
