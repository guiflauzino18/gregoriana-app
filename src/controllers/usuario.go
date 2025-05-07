package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gregorian-app/src/config"
	"gregorian-app/src/dto"
	request "gregorian-app/src/requests"
	"gregorian-app/src/respostas"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CadastrarUsuario(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("%s/api/admin/usuario/cadastro", config.APIURL)

	var usuarioCadastroDTo dto.UsuarioCadastroDTO
	erro := json.NewDecoder(r.Body).Decode(&usuarioCadastroDTo)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	dados, erro := json.Marshal(usuarioCadastroDTo)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	response, erro := request.RequestComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(dados))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

func BuscaUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, erro := strconv.ParseInt(parametros["id"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/api/admin/usuario?id=%d", config.APIURL, id)
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

	var Usuario dto.UsuarioResponseDTO
	if erro := json.NewDecoder(response.Body).Decode(&Usuario); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	respostas.JSON(w, response.StatusCode, &Usuario)
}

func EditaUsuario(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api/admin/usuario/edit", config.APIURL)

	fmt.Println()

	//Passa os dados para o DTO
	var usuario dto.UsuarioEditaDTO
	if erro := json.NewDecoder(r.Body).Decode(&usuario); erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		fmt.Printf("Erro ao decodificar dados passado do form: \n%v", erro)
		return
	}

	// Transforma o DTO em json
	dados, erro := json.Marshal(usuario)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		fmt.Printf("Erro ao converter os dados em DTO para json: \n%v", erro)
		return
	}

	//Faz a requisição à API
	response, erro := request.RequestComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(dados))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		fmt.Printf("Erro ao contatar a API: \n%v", erro)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}

func AlterarSenha(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api/admin/usuario/resetsenha", config.APIURL)

	var usuarioDTO dto.UsuarioAlterarSenhaDTO
	if erro := json.NewDecoder(r.Body).Decode(&usuarioDTO); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	dados, erro := json.Marshal(&usuarioDTO)
	if erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	fmt.Println(bytes.NewBuffer(dados))

	response, erro := request.RequestComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(dados))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, http.StatusOK, nil)

}

func ExcluirUsuario(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	url := fmt.Sprintf("%s/api/admin/usuario/exclui/%s", config.APIURL, id)

	response, erro := request.RequestComAutenticacao(r, http.MethodDelete, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, http.StatusOK, nil)
}
