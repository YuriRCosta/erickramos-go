package rotas

import (
	"erickramos-go/src/controllers"
	"net/http"
)

var rotasCabecotes = []Rota{
	{
		URI:                "/cabecotes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarCabecote,
		RequerAutenticacao: true,
	},
	{
		URI:                "/cabecotes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarCabecotes,
		RequerAutenticacao: true,
	},
	{
		URI:                "/cabecotes/{cabecoteID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarCabecotePorID,
		RequerAutenticacao: true,
	},
	{
		URI:                "/cabecotes/{cabecoteID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarCabecote,
		RequerAutenticacao: true,
	},
	{
		URI:                "/cabecotes/{cabecoteID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarCabecote,
		RequerAutenticacao: true,
	},
	{
		URI:                "/cabecotes/nome/{nomeCabecote}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarCabecotePorNome,
		RequerAutenticacao: true,
	},
}
