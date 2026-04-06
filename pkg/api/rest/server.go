package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server handles the REST API for VaultMesh.
type Server struct {
	router     *gin.Engine
	httpServer *http.Server
}

// NewServer creates a new REST API server.
func NewServer(port int) *Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	return &Server{
		router: r,
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: r,
		},
	}
}

// Start begins listening for API requests.
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully stops the API server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
