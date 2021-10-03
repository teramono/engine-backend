package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login ...
func (server *BackendServer) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ //
		"message": "Logging user in...",
	})
}

// Run ...
func (server *BackendServer) Run(ctx *gin.Context) {
	// TODO: Return response asap while goroutine processes script
	ctx.JSON(http.StatusOK, gin.H{ //
		"message": "Script running...",
	})
}
