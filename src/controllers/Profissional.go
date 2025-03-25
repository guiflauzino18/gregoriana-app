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
