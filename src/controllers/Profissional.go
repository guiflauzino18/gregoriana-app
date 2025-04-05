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

	"github.com/gorilla/mux"
)

func CadastrarProfissional(w http.ResponseWriter, r *http.Request) {

	var profissional dto.ProfissionalCadastroDTO
	if erro := json.NewDecoder(r.Body).Decode(&profissional); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	dados, erro := json.Marshal(profissional)
	if erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/api/admin/profissional/cadastro", config.APIURL)

	fmt.Println(bytes.NewBuffer(dados))

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

func BuscaProfissional(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := fmt.Sprintf("%s/api/admin/profissional/findbyid?id=%s", config.APIURL, params["id"])

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

	var profissional dto.ProfissionalResponseDTO
	if erro := json.NewDecoder(response.Body).Decode(&profissional); erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	respostas.JSON(w, http.StatusOK, profissional)

}

func EditaProfissional(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api/admin/profissional/edit", config.APIURL)

	var profissional dto.ProfissionalEditaDTO
	if erro := json.NewDecoder(r.Body).Decode(&profissional); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	dados, erro := json.Marshal(profissional)
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

func DeletaProfissional(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	url := fmt.Sprintf("%s/api/admin/profissional/delete/%s", config.APIURL, param["id"])

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

func ListaProfissional(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api/admin/profissionais", config.APIURL)

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

	var profissional []dto.ProfissionalListNomeDTO
	if erro := json.NewDecoder(response.Body).Decode(&profissional); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	respostas.JSON(w, http.StatusOK, profissional)
}
