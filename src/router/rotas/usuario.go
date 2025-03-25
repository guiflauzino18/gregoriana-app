package rotas

import (
	"gregorian-app/src/controllers"
	"net/http"
)

var rotaCadastroUsuario = Rota{
	URI:                "/usuario",
	Metodo:             http.MethodPost,
	Funcao:             controllers.CadastrarUsuario,
	RequerAutenticacao: true,
}

var rotaBuscaUsuario = Rota{
	URI:                "/usuario/{id}",
	Metodo:             http.MethodGet,
	Funcao:             controllers.BuscaUsuario,
	RequerAutenticacao: true,
}

var rotaEditaUsuario = Rota{
	URI:                "/usuario",
	Metodo:             http.MethodPut,
	Funcao:             controllers.EditaUsuario,
	RequerAutenticacao: true,
}

var rotaLogout = Rota{
	URI:                "/logout",
	Metodo:             http.MethodGet,
	Funcao:             controllers.FazerLogout,
	RequerAutenticacao: true,
}

var rotaSenha = Rota{
	URI:                "/senha",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarSenha,
	RequerAutenticacao: true,
}

var rotaAlterarSenha = Rota{
	URI:                "/senha",
	Metodo:             http.MethodPost,
	Funcao:             controllers.AlterarSenha,
	RequerAutenticacao: true,
}

var rotaExluiUsuario = Rota{
	URI:                "/usuario/{id}",
	Metodo:             http.MethodDelete,
	Funcao:             controllers.ExcluirUsuario,
	RequerAutenticacao: true,
}
