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

var rotaBuscaProfissional = Rota{
	URI:                "/profissional/{id}",
	Metodo:             http.MethodGet,
	Funcao:             controllers.BuscaProfissional,
	RequerAutenticacao: true,
}

var rotaEditaProfissional = Rota{
	URI:                "/profissional",
	Metodo:             http.MethodPut,
	Funcao:             controllers.EditaProfissional,
	RequerAutenticacao: true,
}

var rotaDeletaProfissional = Rota{
	URI:                "/profissional/{id}",
	Metodo:             http.MethodDelete,
	Funcao:             controllers.DeletaProfissional,
	RequerAutenticacao: true,
}

var rotaListaProfissional = Rota{
	URI:                "/profissionais",
	Metodo:             http.MethodGet,
	Funcao:             controllers.ListaProfissional,
	RequerAutenticacao: true,
}
