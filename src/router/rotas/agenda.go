package rotas

import (
	"gregorian-app/src/controllers"
	"net/http"
)

var rotaAgenda = Rota{
	URI:                "/agenda",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregaAgenda,
	RequerAutenticacao: true,
}

var rotaCadastroAgenda = Rota{
	URI:                "/agenda",
	Metodo:             http.MethodPost,
	Funcao:             controllers.CadastraAgenda,
	RequerAutenticacao: true,
}

var rotaDeletaAgenda = Rota{
	URI:                "/agenda/{id}",
	Metodo:             http.MethodDelete,
	Funcao:             controllers.DeletaAgenda,
	RequerAutenticacao: true,
}

var rotaConfiguraAgenda = Rota{
	URI:                "/agenda/configure",
	Metodo:             http.MethodPut,
	Funcao:             controllers.ConfiguraAgenda,
	RequerAutenticacao: true,
}

var rotaBuscaAgenda = Rota{
	URI:                "/agenda/{id}",
	Metodo:             http.MethodGet,
	Funcao:             controllers.BuscaAgenda,
	RequerAutenticacao: true,
}

var rotaBuscaHorasDoDia = Rota{
	URI:                "/horas/{idDia}",
	Metodo:             http.MethodGet,
	Funcao:             controllers.BuscaHorasaDoDia,
	RequerAutenticacao: true,
}

var rotaStatusHora = []Rota{
	{
		URI:                "/status/hora",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CadastroStatusHora,
		RequerAutenticacao: true,
	},
	{
		URI:                "/agenda/horas/status",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscaStatusHoras,
		RequerAutenticacao: true,
	},
	{
		URI:                "/agenda/hora/status",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AlteraStatusHoras,
		RequerAutenticacao: true,
	},
}
