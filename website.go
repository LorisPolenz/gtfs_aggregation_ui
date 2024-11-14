package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func load_index(c *gin.Context) {
	// Read the HTML content from the file
	htmlContent, err := os.ReadFile("ui/index.html")
	if err != nil {
		log.Printf("Failed to read HTML file: %v", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Serve the content as HTML
	c.Data(http.StatusOK, "text/html; charset=utf-8", htmlContent)
}
