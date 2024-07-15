package server

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ntc-goer/parser-exercise/config"
	"net/http"
)

type HTTPServer struct {
	Engine *gin.Engine
	Server *http.Server
}

func NewHTTPServer(engine *gin.Engine, cfg *config.Config) *HTTPServer {
	server := &http.Server{Addr: fmt.Sprintf(":%s", cfg.ServerPort), Handler: engine}
	return &HTTPServer{
		Engine: engine,
		Server: server,
	}
}

func (h *HTTPServer) Start() error {
	if err := h.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
