package api

import (
	"io/ioutil"
	"log"
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

	chain := alice.New(s.timeoutHandler, s.loggingHandler).Then(routes)
	return chain, s
}

func (s Server) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "web/index.html")
}

func (s Server) getLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	challenge := r.URL.Query().Get("login_challenge")
	log.Printf("level=info challenge=%s", challenge)
	redirectLink, err := s.userstate.ValidateLoginChallenge(challenge)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	}
	if redirectLink == "" {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		http.Redirect(w, r, redirectLink, http.StatusFound)
	}
}

func (s Server) postLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("level=error msg=%s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	var l entities.LoginRequest
	if err := l.Unmarshal(b); err == nil {
		if redirectLink, err := s.userstate.Login(l.Username, l.Password, l.Challenge); err == nil {
			if err != nil {
				log.Printf("level=error msg=%s", err)
				w.WriteHeader(http.StatusUnprocessableEntity)
			} else {
				log.Printf("level=info redirect_link=%s", redirectLink)
				// http.Redirect(w, r, redirectLink, http.StatusFound)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(redirectLink))
			}
		}
	}
	w.WriteHeader(http.StatusUnprocessableEntity)
}

func (s Server) rewriteLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.URL.Path = "/app/--/login"
	http.Redirect(w, r, strings.Replace(r.URL.String(), "--", "#", -1), http.StatusMovedPermanently)
}

func (s Server) rewriteConsent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.URL.Path = "/app/--/consent"
	http.Redirect(w, r, strings.Replace(r.URL.String(), "--", "#", -1), http.StatusMovedPermanently)
}

func (s Server) createRoutes() http.Handler {
	routes := httprouter.New()
	routes.GET("/login", s.rewriteLogin)
	routes.GET("/consent", s.rewriteConsent)
	routes.GET("/api/login", s.getLogin)
	routes.POST("/api/login", s.postLogin)
	routes.ServeFiles("/app/*filepath", http.Dir("web"))
	return routes
}
