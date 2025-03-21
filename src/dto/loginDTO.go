package dto

// LoginDTO é o struct para fazer login
type LoginRequestDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	Token string `json:"token"`
}

type UsuarioResponseDTO struct {
	Id              int        `json:"id"`
	Nome            string     `json:"nome"`
	Sobrenome       string     `json:"sobrenome"`
	Nascimento      string     `json:"nascimento"`
	Telefone        string     `json:"telefone"`
	Email           string     `json:"email"`
	Login           string     `json:"login"`
	Endereco        string     `json:"endereco"`
	AlteraNextLogon bool       `json:"alteraNextLogon"`
	Empresa         EmpresaDTO `json:"empresa"`
}
