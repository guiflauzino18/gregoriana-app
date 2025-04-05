package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gregorian-app/src/config"
	"gregorian-app/src/cookies"
	"gregorian-app/src/dto"
	"gregorian-app/src/respostas"
	"net/http"
)

// FazerLogin utiliza o usuario e senha para autenticar na aplicação
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//FAz o decode do Json e joga os valores dentro do loginRequestDTO
	var loginRequestDTO dto.LoginRequestDTO
	err := json.NewDecoder(r.Body).Decode(&loginRequestDTO)
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	//Cria uma json com loginRequestDTO
	usuario, err := json.Marshal(loginRequestDTO)
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	//Faz requisição à APi
	url := fmt.Sprintf("%s/api/login", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario)) //bytes.NewBuffer(Usuario) converte usuasrio de bytes para json
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//Fecha o body quando finalizar
	defer response.Body.Close()

	//Se statuscode foi maior que 400 retorna erro
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	//Cria o loginResponseDTO com o token recebido
	var loginResponseDTO dto.LoginResponseDTO
	if erro := json.NewDecoder(response.Body).Decode(&loginResponseDTO); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if erro = cookies.Salvar(w, loginResponseDTO.Token, string(loginResponseDTO.ID)); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	respostas.JSON(w, http.StatusOK, nil)

}

func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", 302)
}
