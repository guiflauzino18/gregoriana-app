package rotas

import (
	"gregorian-app/src/controllers"
	"net/http"
)

var rotahome = Rota{
	URI:                "/home",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarHome,
	RequerAutenticacao: true,
}
