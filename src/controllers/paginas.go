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

	if usuario.AlteraNextLogon {
		http.Redirect(w, r, "/senha", 302)

	} else {

		config.Navegacao = []string{}
		config.Navegacao = criaNavegacao("home")

		utils.ExecutarTemplate(w, "index.html", struct {
			Usuario   *dto.UsuarioResponseDTO
			URL       string
			Navegacao []string
		}{
			Usuario:   &usuario,
			URL:       "/home",
			Navegacao: config.Navegacao,
		})
	}
}

// CarregarConfiguracao carrega tela de Configuracao
func CarregarConfiguracao(w http.ResponseWriter, r *http.Request) {

	config.Navegacao = criaNavegacao("configuracao")

	utils.ExecutarTemplate(w, "configuracao.html", struct {
		Usuario   *dto.UsuarioResponseDTO
		URL       string
		Navegacao []string
	}{
		Usuario:   &usuario,
		URL:       "/configuracao",
		Navegacao: config.Navegacao,
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

	existe := false

	for _, pagina := range config.Navegacao {
		if pagina == "usuarios" {
			existe = true
		}

	}

	if !existe {
		config.Navegacao = append(config.Navegacao, "usuarios")
	}

	utils.ExecutarTemplate(w, "usuarios.html", struct {
		Usuarios  *dto.Pageable[dto.UsuariosResponseDTO]
		Usuario   *dto.UsuarioResponseDTO
		URL       string
		Navegacao []string
	}{
		Usuarios:  &usuarios,
		Usuario:   &usuario,
		URL:       "/usuarios",
		Navegacao: config.Navegacao,
	})
}

func CarregarSenha(w http.ResponseWriter, r *http.Request) {
	//Tras dados do usuário logado
	url := fmt.Sprintf("%s/api/index", config.APIURL)
	response, err := request.RequestComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		http.Redirect(w, r, "/login", 302)
		return
	}

	var UsuarioLogado dto.UsuarioResponseDTO
	if err = json.NewDecoder(response.Body).Decode(&UsuarioLogado); err != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: err.Error()})
		http.Redirect(w, r, "/login", 302)
		return
	}

	utils.ExecutarTemplate(w, "senha.html", struct {
		Usuario *dto.UsuarioResponseDTO
		URL     string
	}{
		Usuario: &UsuarioLogado,
		URL:     "/senha",
	})
}

func CarregarProfissional(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api/admin/profissional/list", config.APIURL)

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

	var profissional dto.Pageable[dto.ProfissionalResponseDTO]
	if erro := json.Unmarshal(body, &profissional); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	config.Navegacao = criaNavegacao("profissional")

	utils.ExecutarTemplate(w, "profissional.html", struct {
		Profissional *dto.Pageable[dto.ProfissionalResponseDTO]
		Usuario      *dto.UsuarioResponseDTO
		URL          string
		Navegacao    []string
	}{
		Profissional: &profissional,
		Usuario:      &usuario,
		URL:          "/profissional",
		Navegacao:    config.Navegacao,
	})

}

func CarregaAgenda(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/api/admin/agenda/list", config.APIURL)

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

	var agendas dto.Pageable[dto.AgendaResponseDTO]

	if erro := json.NewDecoder(response.Body).Decode(&agendas); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	fmt.Println(agendas)

	config.Navegacao = criaNavegacao("agenda")

	utils.ExecutarTemplate(w, "agenda.html", struct {
		Agendas   *dto.Pageable[dto.AgendaResponseDTO]
		Usuario   *dto.UsuarioResponseDTO
		URL       string
		Navegacao []string
	}{
		Agendas:   &agendas,
		Usuario:   &usuario,
		URL:       "/agenda",
		Navegacao: config.Navegacao,
	})

}

func CarregaPerfil(w http.ResponseWriter, r *http.Request) {

	utils.ExecutarTemplate(w, "perfil.html", nil)
}

func criaNavegacao(url string) []string {
	existePaginaNoSlice := false
	var novaNavegacao []string

	for _, pagina := range config.Navegacao {

		fmt.Println(pagina)

		novaNavegacao = append(novaNavegacao, pagina)

		if pagina == url {
			config.Navegacao = novaNavegacao
			existePaginaNoSlice = true
			break
		}

	}

	if !existePaginaNoSlice {
		novaNavegacao = append(novaNavegacao, url)
	}

	return novaNavegacao
}
