package main

import (
	"log"
	"net/http"

	"github.com/openpsd/auth-server/providers/api"
	"github.com/openpsd/auth-server/providers/config"
)

func main() {
	conf := config.LoadConfig()
	log.Println("OpenPSD auth server")
	authServer, _ := api.NewServer(conf)
	http.ListenAndServe(":8000", authServer)
}
