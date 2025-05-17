package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gregorian-app/src/config"
	"gregorian-app/src/dto"
	request "gregorian-app/src/requests"
	"gregorian-app/src/respostas"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CadastraAgenda(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api/admin/agenda/cadastro", config.APIURL)

	var agenda dto.AgendaCadastroDTO
	if erro := json.NewDecoder(r.Body).Decode(&agenda); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	dados, erro := json.Marshal(agenda)
	if erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
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

	respostas.JSON(w, http.StatusOK, nil)
}

func DeletaAgenda(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	url := fmt.Sprintf("%s/api/admin/agenda/delete/%d", config.APIURL, id)
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

func ConfiguraAgenda(w http.ResponseWriter, r *http.Request) {

	var agenda dto.AgendaConfiguraDTO

	if erro := json.NewDecoder(r.Body).Decode(&agenda); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	dados, erro := json.Marshal(agenda)
	if erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/api/admin/agenda/config", config.APIURL)

	response, erro := request.RequestComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(dados))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	respostas.JSON(w, http.StatusOK, nil)
}

func BuscaAgenda(w http.ResponseWriter, r *http.Request) {
	// Pega o id da URL
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	url := fmt.Sprintf("%s/api/admin/agenda/%d", config.APIURL, id)

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

	var agendaDTO dto.AgendaResponseDTO
	if erro := json.NewDecoder(response.Body).Decode(&agendaDTO); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	respostas.JSON(w, http.StatusOK, agendaDTO)
}

func BuscaHorasaDoDia(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idDia := params["idDia"]

	url := fmt.Sprintf("%s/api/admin/agenda/horas?id=%s", config.APIURL, idDia)

	response, erro := request.RequestComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
	}

	var horas []dto.HorasRequestDTO

	if erro := json.NewDecoder(response.Body).Decode(&horas); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		fmt.Println(erro)
		return
	}

	respostas.JSON(w, http.StatusOK, horas)

}

func CadastroStatusHora(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("%s/api/admin/agenda/horas/status", config.APIURL)

	// dados, erro := json.Marshal(r.Body)
	// if erro != nil {
	// 	respostas.JSON(w, http.StatusUnprocessableEntity, erro)
	// 	return
	// }

	body, _ := io.ReadAll(r.Body)

	response, erro := request.RequestComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(body))
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

func BuscaStatusHoras(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("%s/api/admin/agenda/horas/status", config.APIURL)

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

	var status []dto.StatusHoraResponseDTO

	if erro := json.NewDecoder(response.Body).Decode(&status); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: fmt.Errorf("Erro no decoder: %v", erro).Error()})
	}

	respostas.JSON(w, response.StatusCode, status)

}

func AlteraStatusHoras(w http.ResponseWriter, r *http.Request) {

	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: fmt.Errorf("Erro no Unmarshal do body: %v", erro).Error()})
		return
	}

	url := fmt.Sprintf("%s/api/admin/agenda/edit/horas", config.APIURL)

	response, erro := request.RequestComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(body))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: fmt.Errorf("Erro no request: %v", erro).Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		retorno, _ := io.ReadAll(response.Body)
		respostas.JSON(w, response.StatusCode, respostas.ErroAPI{Erro: fmt.Sprintf("%v", bytes.NewBuffer(retorno))})
		return
	}

	respostas.JSON(w, response.StatusCode, nil)

}
