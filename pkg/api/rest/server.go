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

// NewServer creates a new REST API server and registers routes.
func NewServer(port int) *Server {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	s := &Server{
		router: r,
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: r,
		},
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.POST("/upload", s.UploadHandler)
	s.router.GET("/download/:cid", s.DownloadHandler)
	s.router.GET("/files", s.ListFilesHandler)
	s.router.GET("/search", s.SearchHandler)
	s.router.DELETE("/files/:cid", s.DeleteFileHandler)
}

// Start begins listening for API requests.
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully stops the API server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
