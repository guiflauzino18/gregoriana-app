package rotas

import (
	"gregorian-app/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Cofigurar adiciona todas as rotas e suas configurações
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotahome)
	rotas = append(rotas, rotaconfiguracao)
	rotas = append(rotas, rotausuarios)
	rotas = append(rotas, rotaCadastroUsuario)
	rotas = append(rotas, rotaBuscaUsuario)
	rotas = append(rotas, rotaEditaUsuario)
	rotas = append(rotas, rotaLogout)
	rotas = append(rotas, rotaSenha)
	rotas = append(rotas, rotaAlterarSenha)
	rotas = append(rotas, rotaExluiUsuario)
	rotas = append(rotas, rotaProfissional)
	rotas = append(rotas, rotaCadastroProfissional)
	rotas = append(rotas, rotaBuscaProfissional)
	rotas = append(rotas, rotaEditaProfissional)
	rotas = append(rotas, rotaDeletaProfissional)
	rotas = append(rotas, rotaListaProfissional)
	rotas = append(rotas, rotaAgenda)
	rotas = append(rotas, rotaCadastroAgenda)
	rotas = append(rotas, rotaDeletaAgenda)
	rotas = append(rotas, rotaPerfil)
	rotas = append(rotas, rotaBuscaAgenda)
	rotas = append(rotas, rotaConfiguraAgenda)
	rotas = append(rotas, rotaBuscaHorasDoDia)
	rotas = append(rotas, rotaStatusHora...)

	for _, rota := range rotas {
		if rota.RequerAutenticacao { //Se Rota requer Autenticação é chamado um Middleware passando primeiro a autenticação e depois chamando a função
			router.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}

	fileserver := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileserver))
	return router
}
