package model

import (
	"errors"
)

//Pessoa é uma entidade de uma pessoa
type Pessoa struct {
	nome     string
	idade    int
	telefone string
}

//SetNome é o set do nome da estrutura de pessoa
func (p *Pessoa) SetNome(nome string) {
	p.nome = nome
}

//SetIdade é o set da idade da estrutura de pessoa
func (p *Pessoa) SetIdade(idade int) (err error) {
	err = nil
	if idade < 18 {
		err = ErrIdadeInvalida
		return
	}
	p.idade = idade
	return
}

//SetTelefone é o set do telefone da estrutura de pessoa
func (p *Pessoa) SetTelefone(telefone string) {
	p.telefone = telefone
}

//ErrIdadeInvalida é retornado quando a idade é menor que 18 anos
var ErrIdadeInvalida = errors.New("O valor da idade está inválido")
