package webserver

import (
	"context"
	"desafio_mercado_livre/src/internal/infra/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Path        string
	Method      string
	HandlerFunc gin.HandlerFunc
}

type WebServer struct {
	Logger        logger.Logger
	Router        *gin.Engine
	Handlers      []Handler
	WebServerPort string
}

func NewWebServer(logger logger.Logger, router *gin.Engine, port string) *WebServer {
	return &WebServer{
		Logger:        logger,
		Router:        router,
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
	for _, handler := range w.Handlers {
		w.Router.Handle(handler.Method, handler.Path, handler.HandlerFunc)
	}

	serverLog := fmt.Sprintf("Server listening on port %s", w.WebServerPort)
	w.Logger.Info(serverLog)

	srv := &http.Server{
		Addr:    w.WebServerPort,
		Handler: w.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			w.Logger.Error("Erro ao iniciar o servidor", "error", err)
		}
	}()

	// Captura sinais de interrupção
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	w.Logger.Info("Desligando o servidor...")

	// Timeout para shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		w.Logger.Error("Erro no shutdown do servidor", "error", err)
	}
	w.Logger.Info("Servidor finalizado com sucesso")
}
