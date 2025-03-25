package dto

type ProfissionalResponseDTO struct {
	Id       int                `json:"id"`
	Titulo   string             `json:"titulo"`
	Registro string             `json:"registro"`
	Usuario  UsuarioResponseDTO `json:"usuario"`
}

type ProfissionalCadastroDTO struct {
	Titulo   string `json:"titulo"`
	Registro string `json:"registro"`
	Login    string `json:"login"`
}
