package webserver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]map[string]http.HandlerFunc), // httpmethod -> path -> handlerFunc
		WebServerPort: serverPort,
	}

}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	if s.Handlers[method] == nil {
		s.Handlers[method] = make(map[string]http.HandlerFunc)
	}

	s.Handlers[method][path] = handler
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for httpMethod, pathToHandler := range s.Handlers {
		for path, handler := range pathToHandler {
			s.Router.MethodFunc(httpMethod, path, handler)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
