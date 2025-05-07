package rotas

import (
	"gregorian-app/src/controllers"
	"net/http"
)

var rotaconfiguracao = Rota{
	URI:                "/configuracao",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarConfiguracao,
	RequerAutenticacao: true,
}

var rotausuarios = Rota{
	URI:                "/usuarios",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarUsuarios,
	RequerAutenticacao: true,
}
