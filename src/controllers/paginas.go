package controllers

import (
	"encoding/json"
	"fmt"
	"gregorian-app/src/config"
	"gregorian-app/src/cookies"
	"gregorian-app/src/dto"
	request "gregorian-app/src/requests"
	"gregorian-app/src/respostas"
	"gregorian-app/src/utils"
	"io"
	"net/http"
)

var usuario dto.UsuarioResponseDTO

// CarregarTelaDeLogin carrega tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	//Se carergar cookie é redirecionado para tela Home
	cookie, _ := cookies.Ler(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarHome carrega página principal
func CarregarHome(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api/index", config.APIURL)

	response, err := request.RequestComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	if err = json.NewDecoder(response.Body).Decode(&usuario); err != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	utils.ExecutarTemplate(w, "index.html", struct {
		Usuario *dto.UsuarioResponseDTO
		URL     string
	}{
		Usuario: &usuario,
		URL:     "/home",
	})

}

// CarregarConfiguracao carrega tela de Configuracao
func CarregarConfiguracao(w http.ResponseWriter, r *http.Request) {

	utils.ExecutarTemplate(w, "configuracao.html", struct {
		Usuario *dto.UsuarioResponseDTO
		URL     string
	}{
		Usuario: &usuario,
		URL:     "/configuracao",
	})
}

func CarregarUsuarios(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api/admin/usuario/list", config.APIURL)

	response, erro := request.RequestComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	//Pega conteudo do Body
	body, erro := io.ReadAll(response.Body)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// Joga conteudo do Body dentro do struct
	var usuarios dto.Pageable[dto.UsuariosResponseDTO]
	if err := json.Unmarshal(body, &usuarios); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}

	utils.ExecutarTemplate(w, "usuarios.html", struct {
		Usuarios *dto.Pageable[dto.UsuariosResponseDTO]
		Usuario  *dto.UsuarioResponseDTO
		URL      string
	}{
		Usuarios: &usuarios,
		Usuario:  &usuario,
		URL:      "/usuarios",
	})
}
