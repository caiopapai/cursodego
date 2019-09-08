package model

// Galinha é a representação de um tipo de ave
type Galinha interface {
	Cacareja() string
}

// Pato é a representação de um tipo de ave
type Pato interface {
	Quack() string
}

//Ave é um tipo do reino animal
type Ave struct {
	Nome string
}

//Cacareja representa um som feito por uma galinha?
func (a Ave) Cacareja() string {
	return "Cócóricó..."
}

//Quack representa um som feito por uma galinha?
func (a Ave) Quack() string {
	return "Quack Quack"
}
