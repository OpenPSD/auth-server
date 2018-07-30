package api

import (
	"io/ioutil"
	"net/http"
	"strings"

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
	http.ServeFile(w, r, "web/index.html")
}

func (s Server) getLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	challenge := r.URL.Query().Get("login_challenge")
	if redirectLink, err := s.userstate.ValidateLoginChallenge(challenge); err == nil {
		if redirectLink == "" {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			http.Redirect(w, r, redirectLink, http.StatusMovedPermanently)
		}
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func (s Server) postLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	challenge := r.URL.Query().Get("login_challenge")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	var l entities.LoginRequest
	if err := l.Unmarshal(b); err == nil {
		if redirectLink, err := s.userstate.Login(l.Username, l.Password, challenge); err == nil {
			http.Redirect(w, r, redirectLink, http.StatusMovedPermanently)
		}

	}
	w.WriteHeader(http.StatusUnprocessableEntity)
}

func (s Server) rewriteLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.URL.Path = "/app/--/login--"
	http.Redirect(w, r, strings.Replace(r.URL.String(), "--", "#", -1), http.StatusMovedPermanently)
}

func (s Server) rewriteConsent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.URL.Path = "/app/--/consent--"
	http.Redirect(w, r, strings.Replace(r.URL.String(), "--", "#", -1), http.StatusMovedPermanently)
}

func (s Server) createRoutes() http.Handler {
	routes := httprouter.New()
	routes.GET("/login", s.rewriteLogin)
	routes.GET("/consent", s.rewriteConsent)
	routes.GET("/api/validate", s.getLogin)
	routes.POST("/api/login", s.postLogin)
	routes.ServeFiles("/app/*filepath", http.Dir("web"))
	return routes
}
