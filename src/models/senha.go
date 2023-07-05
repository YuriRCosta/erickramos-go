package models

// Senha representa a senha de um usuário
type Senha struct {
	Nova   string `json:"nova"`
	Antiga string `json:"antiga"`
}
