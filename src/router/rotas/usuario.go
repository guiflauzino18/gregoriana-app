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
