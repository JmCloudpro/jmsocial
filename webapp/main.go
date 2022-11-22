package main

import (
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()

	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
	log.Fatal(http.ListenAndServe(":8080", r))
} //
