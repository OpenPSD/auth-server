package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/openpsd/auth-server/server/entities"
	"github.com/openpsd/auth-server/server/providers/api"
	"github.com/openpsd/auth-server/server/providers/log"
	"github.com/openpsd/auth-server/server/providers/oauth"
	"github.com/openpsd/auth-server/server/providers/userstore"
	"github.com/openpsd/auth-server/server/usecases"
)

func init() {
	log.InitLog(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}

func main() {
	config := entities.NewConfig()
	log.Info.Println("OpenPSD auth server")
	userstore := userstore.NewMemUserStore()
	oauthclient, err := oauth.NewHydraClient(config.HydraURL)
	log.Info.Printf("hydraURL=%s port=%s", oauthclient.HydraSDK.Configuration.EndpointURL, config.Port)
	if err != nil {
		log.Error.Println("msg='able to create oauth client'")
	}
	userstate := usecases.NewUserstate(userstore, oauthclient)
	userstate.CreateUser("admin", "admin@openpsd.com", "admin")
	log.Info.Println("msg='created demo admin user'")
	authServer, _ := api.NewServer(userstate)
	http.ListenAndServe(config.Port, authServer)
}
