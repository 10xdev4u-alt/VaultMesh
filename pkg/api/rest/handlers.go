package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
