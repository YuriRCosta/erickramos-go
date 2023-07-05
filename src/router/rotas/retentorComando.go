package rotas

import (
	"erickramos-go/src/controllers"
	"net/http"
)

var rotasRetentorComando = []Rota{
	{
		URI:                "/retentorComando",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarRetentorComando,
		RequerAutenticacao: true,
	},
	{
		URI:                "/retentorComando",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarRetentoresComando,
		RequerAutenticacao: true,
	},
	{
		URI:                "/retentorComando/{retentorComandoID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarRetentorComandoPorID,
		RequerAutenticacao: true,
	},
	{
		URI:                "/retentorComando/{retentorComandoID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarRetentorComando,
		RequerAutenticacao: true,
	},
	{
		URI:                "/retentorComando/{retentorComandoID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarRetentorComando,
		RequerAutenticacao: true,
	},
	{
		URI:                "/retentorComando/nome/{nomeRetentorComando}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarRetentoresComandoPorNome,
		RequerAutenticacao: true,
	},
}
