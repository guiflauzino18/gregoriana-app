package dto

type AgendaResponseDTO struct {
	Id               int    `json:"id"`
	Nome             string `json:"nome"`
	EmpresaNome      string `json:"empresaNome"`
	IdProfissional   int    `json:"idProfissional"`
	NomeProfissional string `json:"NomeProfissional"`
	StatusAgenda     struct {
		ID   int    `json:"id"`
		Nome string `json:"nome"`
	} `json:"status"`
	Dias []Dias `json:"dias"`
}

type AgendaCadastroDTO struct {
	Nome         string `json:"nome"`
	Profissional int    `json:"idProfissional"`
}

type AgendaConfiguraDTO struct {
	IdAgenda       int    `json:"idAgenda"`
	IdProfissional int    `json:"idProfissional"`
	Dias           []Dias `json:"dias"`
}

type Dias struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Intervalo int    `json:"intervaloSessaoInMinutes"`
	Duracao   int    `json:"duracaoSessaoInMinutes"`
	Inicio    string `json:"inicio"`
	Fim       string `json:"fim"`
}

type HorasRequestDTO struct {
	ID     int    `json:"id"`
	Inicio string `json:"inicio"`
	Fim    string `json:"fim"`
	Status struct {
		ID   int    `json:"id"`
		Nome string `json:"nome"`
	} `json:"statusHora"`
}

type StatusHoraResponseDTO struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}
