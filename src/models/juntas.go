package models

type Junta struct {
	ID          uint64  `json:"id,omitempty"`
	Cabecotes   string  `json:"cabecotes,omitempty"`
	Preco       float64 `json:"preco,omitempty"`
	Qtd_estoque int     `json:"qtd_estoque,omitempty"`
}
