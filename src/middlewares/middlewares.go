package middlewares

import (
	"erickramos-go/src/auth"
	"erickramos-go/src/respostas"
	"log"
	"net/http"
)

// Logger escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidarToken(r); err != nil {
			respostas.Erro(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
