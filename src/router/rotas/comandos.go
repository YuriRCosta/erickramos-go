package rotas

import (
	"erickramos-go/src/controllers"
	"net/http"
)

var rotasComandos = []Rota{
	{
		URI:                "/comandos",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarComando,
		RequerAutenticacao: true,
	},
	{
		URI:                "/comandos",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarComandos,
		RequerAutenticacao: true,
	},
	{
		URI:                "/comandos/{comandoID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarComandoPorID,
		RequerAutenticacao: true,
	},
	{
		URI:                "/comandos/{comandoID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarComando,
		RequerAutenticacao: true,
	},
	{
		URI:                "/comandos/{comandoID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarComando,
		RequerAutenticacao: true,
	},
	{
		URI:                "/comandos/nome/{nomeComando}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarComandoPorNome,
		RequerAutenticacao: true,
	},
}
