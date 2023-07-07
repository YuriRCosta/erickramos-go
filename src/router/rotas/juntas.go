package rotas

import (
	"erickramos-go/src/controllers"
	"net/http"
)

var rotasJuntas = []Rota{
	{
		URI:                "/juntas",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarJunta,
		RequerAutenticacao: true,
	},
	{
		URI:                "/juntas",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarJuntas,
		RequerAutenticacao: true,
	},
	{
		URI:                "/juntas/{juntaID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarJuntaPorID,
		RequerAutenticacao: true,
	},
	{
		URI:                "/juntas/{juntaID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarJunta,
		RequerAutenticacao: true,
	},
	{
		URI:                "/juntas/{juntaID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarJunta,
		RequerAutenticacao: true,
	},
	{
		URI:                "/juntas/cabecotes/{cabecotesJunta}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarJuntaPorCabecotes,
		RequerAutenticacao: true,
	},
}
