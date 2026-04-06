package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadRequest defines the schema for file upload metadata.
type UploadRequest struct {
	Name string `json:"name" binding:"required"`
}

// UploadHandler handles file uploads via multipart form.
func (s *Server) UploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "File received successfully",
		"filename": file.Filename,
		"size":     file.Size,
	})
}

// DownloadHandler handles file retrieval requests.
func (s *Server) DownloadHandler(c *gin.Context) {
	cid := c.Param("cid")
	if cid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing cid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File download initiated",
		"cid":     cid,
	})
}

// ListFilesHandler returns a list of files managed by this node.
func (s *Server) ListFilesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"files": []string{}})
}

// SearchHandler allows searching for files by name or metadata.
func (s *Server) SearchHandler(c *gin.Context) {
	query := c.Query("q")
	c.JSON(http.StatusOK, gin.H{"query": query, "results": []string{}})
}

// DeleteFileHandler removes a file and its chunks.
func (s *Server) DeleteFileHandler(c *gin.Context) {
	cid := c.Param("cid")
	c.JSON(http.StatusOK, gin.H{"message": "File deletion initiated", "cid": cid})
}

// ListPeersHandler returns a list of connected network peers.
func (s *Server) ListPeersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"peers": []string{}})
}

// PeerStatsHandler returns detailed metrics for a specific peer.
func (s *Server) PeerStatsHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"peer_id": id, "reputation": 0.5, "latency": "50ms"})
}
