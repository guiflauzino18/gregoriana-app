package controllers

import (
	"gregorian-app/src/cookies"
	"gregorian-app/src/utils"
	"net/http"
)

func CarregarHome(w http.ResponseWriter, r *http.Request) {

}

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	//Se carergar cookie é redirecionado para tela Home
	cookie, _ := cookies.Ler(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

	utils.ExecutarTemplate(w, "login.html", nil)
}
