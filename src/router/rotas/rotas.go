package rotas

import (
	"erickramos-go/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

// Configurar vai receber um router e retornar um router com as rotas configuradas
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasSelos...)
	rotas = append(rotas, rotasRetentorComando...)
	rotas = append(rotas, rotasRetentorValvula...)
	rotas = append(rotas, rotasValvulas...)
	rotas = append(rotas, rotasComandos...)
	rotas = append(rotas, rotasJuntas...)
	rotas = append(rotas, rotasCabecotes...)

	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

		r.HandleFunc(rota.URI, OptionsHandler).Methods("OPTIONS")
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}

func OptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	w.WriteHeader(http.StatusOK)
}
