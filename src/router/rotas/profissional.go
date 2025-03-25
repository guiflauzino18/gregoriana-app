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

var rotaCadastroProfissional = Rota{
	URI:                "/profissional",
	Metodo:             http.MethodPost,
	Funcao:             controllers.CadastrarProfissional,
	RequerAutenticacao: true,
}
