package main

import (
	"log"
	"net/http"

	"github.com/openpsd/auth-server/providers/api"
	"github.com/openpsd/auth-server/providers/config"
	"github.com/openpsd/auth-server/providers/userstore"
	"github.com/openpsd/auth-server/usecases"
)

func main() {
	conf := config.LoadConfig()
	log.Println("OpenPSD auth server")
	userstore := userstore.NewMemUserStore()
	userstate := usecases.NewUserstate(userstore)
	userstate.CreateUser("admin", "admin@openpsd.com", "admin")
	authServer, _ := api.NewServer(conf, userstate)
	http.ListenAndServe(":8000", authServer)
}
