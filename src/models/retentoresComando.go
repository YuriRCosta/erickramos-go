package models

type RetentorComando struct {
	ID          uint64  `json:"id,omitempty"`
	Nome        string  `json:"nome,omitempty"`
	Preco       float64 `json:"preco,omitempty"`
	Qtd_estoque int     `json:"qtd_estoque,omitempty"`
}
