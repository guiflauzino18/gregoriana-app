package dto

type ProfissionalResponseDTO struct {
	Id        int    `json:"id"`
	Titulo    string `json:"titulo"`
	Registro  string `json:"registro"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Login     string `json:"login"`
	Empresa   string `json:"empresaNome"`
}

type ProfissionalCadastroDTO struct {
	Titulo   string `json:"titulo"`
	Registro string `json:"registro"`
	Login    string `json:"login"`
}

type ProfissionalEditaDTO struct {
	Id       string `json:"id"`
	Titulo   string `json:"titulo"`
	Registro string `json:"registro"`
	Login    string `json:"login"`
	AgendaId string `json:"agendaId"`
}

type ProfissionalListNomeDTO struct {
	Nome string `json:"nome"`
	Id   int    `json:"id"`
}
