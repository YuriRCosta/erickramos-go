package rotas

import (
	"erickramos-go/src/controllers"
	"net/http"
)

var rotasRetentorValvula = []Rota{
	{
		URI:                "/retentorValvula",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarRetentorValvula,
		RequerAutenticacao: true,
	},
	{
		URI:                "/retentorValvula",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarRetentoresValvula,
		RequerAutenticacao: true,
	},
	{
		URI:                "/retentorValvula/{retentorValvulaID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarRetentorValvulaPorID,
		RequerAutenticacao: true,
	},
	{
		URI:                "/retentorValvula/{retentorValvulaID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarRetentorValvula,
		RequerAutenticacao: true,
	},
	{
		URI:                "/retentorValvula/{retentorValvulaID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarRetentorValvula,
		RequerAutenticacao: true,
	},
	{
		URI:                "/retentorValvula/nome/{nomeRetentorValvula}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarRetentoresValvulaPorNome,
		RequerAutenticacao: true,
	},
}
