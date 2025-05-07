package request

import (
	"gregorian-app/src/cookies"
	"io"
	"net/http"
)

// RequestComAutenticacao faz requisições à API colocando o token no cabeçalho
func RequestComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	//Cria o request com os dados
	request, err := http.NewRequest(metodo, url, dados)
	if err != nil {
		return nil, err
	}

	//Le o token do cookie e adiciona no Header do request
	cookie, _ := cookies.Ler(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])
	request.Header.Add("Content-Type", "application/json")

	//Cria um client para fazer a requisição
	client := &http.Client{}

	//Cria um response e armazena o retorno
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	//Retorna o Responde
	return response, nil
}

// FazerRequisicaoComAutenticacao é utilizada para colocar o token na requisição
// func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
// 	request, erro := http.NewRequest(metodo, url, dados)
// 	if erro != nil {
// 		return nil, erro
// 	}

// 	cookie, _ := cookies.Ler(r)
// 	request.Header.Add("Authorization", "Bearer "+cookie["token"])

// 	client := &http.Client{}
// 	response, erro := client.Do(request)
// 	if erro != nil {
// 		return nil, erro
// 	}

// 	return response, nil
// }
