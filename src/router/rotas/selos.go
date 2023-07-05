package rotas

import (
	"erickramos-go/src/controllers"
	"net/http"
)

var rotasSelos = []Rota{
	{
		URI:                "/selos",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarSelo,
		RequerAutenticacao: true,
	},
	{
		URI:                "/selos",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSelos,
		RequerAutenticacao: true,
	},
	{
		URI:                "/selos/{seloID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeloPorID,
		RequerAutenticacao: true,
	},
	{
		URI:                "/selos/{seloID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarSelo,
		RequerAutenticacao: true,
	},
	{
		URI:                "/selos/{seloID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarSelo,
		RequerAutenticacao: true,
	},
	{
		URI:                "/selos/nome/{nomeSelo}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeloPorNome,
		RequerAutenticacao: true,
	},
	{
		URI:                "/selos/medida/{medidaSelo}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeloPorMedida,
		RequerAutenticacao: true,
	},
}
