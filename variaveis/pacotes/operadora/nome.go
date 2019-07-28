package operadora

import (
	"strconv"

	"github.com/caiopapai/variaveis/pacotes/prefixo"
)

//NomeOperadora de Telecom
var NomeOperadora = "Vivo"

//PrefixoCapitalOperadora teste
var PrefixoCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
