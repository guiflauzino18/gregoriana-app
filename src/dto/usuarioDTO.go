package dto

import "time"

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
	Status          string     `json:"status"`
	DataRegistro    time.Time  `json:"dataRegistro"`
	Role            string     `json:"role"`
}

type UsuarioEditaDTO struct {
	Id              string `json:"id"`
	Nome            string `json:"nome"`
	Sobrenome       string `json:"sobrenome"`
	Nascimento      string `json:"nascimento"`
	Telefone        string `json:"telefone"`
	Email           string `json:"email"`
	Endereco        string `json:"endereco"`
	AlteraNextLogon bool   `json:"alteraNextLogon"`
	Status          string `json:"status"`
	Role            string `json:"role"`
}

type UsuarioAlterarSenhaDTO struct {
	Id              string `json:"id"`
	Senha           string `json:"senha"`
	AlteraNextLogon bool   `json:"alteraNextLogon"`
}
