package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Path        string
	Method      string
	HandlerFunc gin.HandlerFunc
}

type WebServer struct {
	Router        *gin.Engine
	Handlers      []Handler
	WebServerPort string
}

func NewWebServer(router *gin.Engine, port string) *WebServer {
	return &WebServer{
		Router:        gin.Default(),
		Handlers:      []Handler{},
		WebServerPort: port,
	}
}

func (w *WebServer) AddHandler(method string, path string, handler gin.HandlerFunc) {
	w.Handlers = append(w.Handlers, Handler{
		Path:        path,
		Method:      method,
		HandlerFunc: handler,
	})
}

func (w *WebServer) Start() {
	//w.Router.Use() add logger
	for _, handler := range w.Handlers {
		w.Router.Handle(handler.Method, handler.Path, handler.HandlerFunc)
	}
	http.ListenAndServe(":"+w.WebServerPort, w.Router)
}
