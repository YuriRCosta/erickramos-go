package rotas

import (
	"erickramos-go/src/controllers"
	"net/http"
)

var rotasRetentorComando = []Rota{
	{
		URI:                "/retentorcomando",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarRetentorComando,
		RequerAutenticacao: true,
	},
	{
		URI:                "/retentorcomando",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarRetentoresComando,
		RequerAutenticacao: true,
	},
	{
		URI:                "/retentorcomando/{retentorcomandoID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarRetentorComandoPorID,
		RequerAutenticacao: true,
	},
	{
		URI:                "/retentorcomando/{retentorcomandoID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarRetentorComando,
		RequerAutenticacao: true,
	},
	{
		URI:                "/retentorcomando/{retentorcomandoID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarRetentorComando,
		RequerAutenticacao: true,
	},
	{
		URI:                "/retentorcomando/nome/{nomeRetentorComando}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarRetentoresComandoPorNome,
		RequerAutenticacao: true,
	},
}
