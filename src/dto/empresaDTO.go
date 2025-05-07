package dto

type EmpresaDTO struct {
	Nome        string `json:"nome"`
	CNPJ        int    `json:"cnpj"`
	Endereco    string `json:"endereco"`
	Telefone    string `json:"telefone"`
	Responsavel string `json:"responsavel"`
}
