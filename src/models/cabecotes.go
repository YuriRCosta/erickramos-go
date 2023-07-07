package models

type Cabecote struct {
	ID           uint64 `json:"id,omitempty"`
	Nome         string `json:"nome,omitempty"`
	Qtd_valvulas int    `json:"qtd_valvulas,omitempty"`
	Material     string `json:"material,omitempty"`
}
