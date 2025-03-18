package main

import (
	"fmt"
	"gregorian-app/src/config"
	"gregorian-app/src/cookies"
	"gregorian-app/src/router"
	"gregorian-app/src/utils"
	"log"
	"net/http"
)

// func init() {
//	Gera o HaskKey e o blockKey usando o pacote securecookie
// 	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(hashKey)

// 	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(blockKey)
// }

func main() {
	//Carrega configurações como variáveis de ambientes
	config.Carregar()

	//Configura cookies
	cookies.Configurar()

	//Carrega os templates
	utils.CarregarTemplates()

	//Carrega rotas configuradas
	r := router.Gerar()

	fmt.Printf("Aplicação iniciada e escutando na porta %s.\nURL da API: %s", config.Porta, config.APIURL)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
