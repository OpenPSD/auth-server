package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/openpsd/auth-server/server/entities"
	"github.com/openpsd/auth-server/server/usecases"
)

// Server holds all http handlers for the PSD2 API
type Server struct {
	userstate *usecases.Userstate
}

// NewServer injects the required dependencies into the PSD2 API server
func NewServer(u *usecases.Userstate) (http.Handler, Server) {
	s := Server{
		userstate: u,
	}

	routes := s.createRoutes()
	chain := alice.New(s.timeoutHandler).Then(routes)
	return chain, s
}

func (s Server) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "OpenPSD auth server is running!\n")
}

func (s Server) postLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	challenge := r.URL.Query().Get("challenge")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	var l entities.LoginRequest
	if err := l.Unmarshal(b); err == nil {
		if redirectLink, err := s.userstate.Login(l.Username, l.Password, challenge); err == nil {
			http.Redirect(w, r, redirectLink, http.StatusMovedPermanently)
		}
		w.WriteHeader(http.StatusUnauthorized)
	}
	w.WriteHeader(http.StatusUnprocessableEntity)
}

func (s Server) createRoutes() http.Handler {
	routes := httprouter.New()
	routes.POST("/login", s.postLogin)
	routes.ServeFiles("/*filepath", http.Dir("web"))
	return routes
}
