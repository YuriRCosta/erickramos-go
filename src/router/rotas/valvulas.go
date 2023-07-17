package rotas

import (
	"erickramos-go/src/controllers"
	"net/http"
)

var rotasValvulas = []Rota{
	{
		URI:                "/valvulas",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarValvula,
		RequerAutenticacao: true,
	},
	{
		URI:                "/valvulas",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarValvulas,
		RequerAutenticacao: true,
	},
	{
		URI:                "/valvulas/pagina/{pagina}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarValvulasPagina,
		RequerAutenticacao: true,
	},
	{
		URI:                "/valvulas/{valvulaID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarValvulaPorID,
		RequerAutenticacao: true,
	},
	{
		URI:                "/valvulas/{valvulaID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarValvula,
		RequerAutenticacao: true,
	},
	{
		URI:                "/valvulas/{valvulaID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarValvula,
		RequerAutenticacao: true,
	},
	{
		URI:                "/valvulas/nome/{nomeValvula}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarValvulaPorNome,
		RequerAutenticacao: true,
	},
	{
		URI:                "/valvulas/tipo/{tipoValvula}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarValvulaPorTipo,
		RequerAutenticacao: true,
	},
	{
		URI:                "/valvulas/adicionar-estoque/{valvulaID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AdicionarEstoqueValvula,
		RequerAutenticacao: true,
	},
}
