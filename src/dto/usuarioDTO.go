package dto

type UsuariosResponseDTO struct {
	Id         int    `json:"id"`
	Nome       string `json:"nome"`
	Sobrenome  string `json:"sobrenome"`
	Nascimento string `json:"nascimento"`
	Telefone   string `json:"telefone"`
	Email      string `json:"email"`
	Login      string `json:"login"`
	Endereco   string `json:"endereco"`
	Role       string `json:"role"`
}

type UsuarioCadastroDTO struct {
	Nome        string `json:"nome"`
	Sobrenome   string `json:"sobrenome"`
	Nascimento  string `json:"nascimento"`
	Telefone    string `json:"telefone"`
	Email       string `json:"email"`
	Login       string `json:"login"`
	Endereco    string `json:"endereco"`
	Role        string `json:"role"`
	Senha       string `json:"senha"`
	AlteraSenha bool   `json:"alteraNextLogon"`
}
