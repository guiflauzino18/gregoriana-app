package dto

// LoginDTO é o struct para fazer login
type LoginRequestDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	Token string `json:"token"`
	ID    int    `json:"id"`
}
