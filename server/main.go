package main

import (
	"log"
	"net/http"

	"github.com/openpsd/auth-server/server/entities"
	"github.com/openpsd/auth-server/server/providers/api"
	"github.com/openpsd/auth-server/server/providers/oauth"
	"github.com/openpsd/auth-server/server/providers/userstore"
	"github.com/openpsd/auth-server/server/usecases"
)

func main() {
	config := entities.NewConfig()
	log.Println("OpenPSD auth server")
	userstore := userstore.NewMemUserStore()
	oauthclient := oauth.NewHydraClient(config.HydraURL)
	userstate := usecases.NewUserstate(userstore, oauthclient)
	userstate.CreateUser("admin", "admin@openpsd.com", "admin")

	authServer, _ := api.NewServer(userstate)
	http.ListenAndServe(config.Port, authServer)
}