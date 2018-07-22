package api

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/openpsd/auth-server/providers/config"
)

// Server holds all http handlers for the PSD2 API
type Server struct {
	config *config.Config
}

// NewServer creates an mocked PSD2 API server
func NewServer(conf *config.Config) (http.Handler, Server) {
	return ServerFactory(conf)
}

// ServerFactory injects the required dependencies into the PSD2 API server
func ServerFactory(conf *config.Config) (http.Handler, Server) {
	s := Server{
		config: conf,
	}

	routes := s.createRoutes()
	chain := alice.New(s.timeoutHandler).Then(routes)
	return chain, s
}

func (s Server) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "OpenPSD auth server is running!\n")
}

func (s Server) getLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	invalid, err := strconv.ParseBool(r.URL.Query().Get("invalid"))
	if err != nil {
		invalid = false
	}
	funcMap := template.FuncMap{
		"lower": func(in string) string { return strings.ToLower(in) },
	}
	t := template.Must(template.New("login.html").Funcs(funcMap).ParseFiles("web/templates/login.html", fmt.Sprintf("%s/header.html", s.config.TemplatesPath), "web/templates/footer.html"))

	data := struct {
		Issuer         string
		URL            string
		PostURL        string
		UsernamePrompt string
		Username       string
		Invalid        bool
	}{"ASPSP", s.config.URL, fmt.Sprintf("%s/login", s.config.URL), "Username", "user1", invalid}
	t.ExecuteTemplate(w, "login.html", data)
}

func (s Server) postLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/login?invalid=true", http.StatusSeeOther)
}

func (s Server) createRoutes() http.Handler {
	routes := httprouter.New()
	routes.GET("/", s.index)
	routes.GET("/login", s.getLogin)
	routes.POST("/login", s.postLogin)
	routes.ServeFiles("/static/*filepath", http.Dir("web/static"))
	routes.ServeFiles("/theme/*filepath", http.Dir(fmt.Sprintf("web/themes/%s", s.config.Theme)))
	return routes
}
