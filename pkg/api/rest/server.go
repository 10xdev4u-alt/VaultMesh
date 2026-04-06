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
	s.router.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	authorized := s.router.Group("/")
	authorized.Use(AuthMiddleware("default-secret-key"))
	{
		authorized.POST("/upload", s.UploadHandler)
		authorized.GET("/download/:cid", s.DownloadHandler)
		authorized.GET("/files", s.ListFilesHandler)
		authorized.GET("/search", s.SearchHandler)
		authorized.DELETE("/files/:cid", s.DeleteFileHandler)
		authorized.GET("/peers", s.ListPeersHandler)
		authorized.GET("/peers/:id/stats", s.PeerStatsHandler)
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
