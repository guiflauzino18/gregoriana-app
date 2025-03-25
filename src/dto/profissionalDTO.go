package dto

type ProfissionalResponseDTO struct {
	Id       int                `json:"id"`
	Titulo   string             `json:"titulo"`
	Registro string             `json:"registro"`
	Usuario  UsuarioResponseDTO `json:"usuario"`
}
