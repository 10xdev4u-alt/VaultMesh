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
