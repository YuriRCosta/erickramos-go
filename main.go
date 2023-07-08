package main

import (
	"erickramos-go/src/config"
	"erickramos-go/src/router"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	config.Carregar()
	r := router.Gerar()

	c := cors.New(cors.Options{
		AllowedMethods:       []string{"GET", "POST", "OPTIONS"},
		AllowedOrigins:       []string{"*"},
		AllowCredentials:     true,
		AllowedHeaders:       []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept"},
		OptionsPassthrough:   true,
		OptionsSuccessStatus: 200,
	})

	handler := c.Handler(r)

	fmt.Printf(config.StringConexaoBanco)
	fmt.Printf("Rodando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), handler))
}
