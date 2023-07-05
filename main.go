package main

import (
	"erickramos-go/src/config"
	"erickramos-go/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.Gerar()

	fmt.Printf(config.StringConexaoBanco)
	fmt.Printf("Rodando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
