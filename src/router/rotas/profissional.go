package rotas

import (
	"gregorian-app/src/controllers"
	"net/http"
)

var rotaProfissional = Rota{
	URI:                "/profissional",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarProfissional,
	RequerAutenticacao: true,
}
