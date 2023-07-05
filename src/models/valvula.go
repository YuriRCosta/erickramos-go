package models

type Valvula struct {
	ID          uint64  `json:"id,omitempty"`
	Codigo      string  `json:"codigo,omitempty"`
	Nome        string  `json:"nome,omitempty"`
	Preco       float64 `json:"preco,omitempty"`
	Qtd_estoque int     `json:"qtd_estoque,omitempty"`
	Tipo        string  `json:"tipo,omitempty"`
}
