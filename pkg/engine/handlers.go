package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login ...
func (server *BackendServer) Login(ctx *gin.Context) {
}

// Run ...
func (server *BackendServer) Run(ctx *gin.Context) {
	// TODO: Will eventually use grpc.
	var body map[string]interface{}

	ctx.BindJSON(&body) // Error is not useful here since body is sometimes allowed not to exist

	ctx.JSON(http.StatusOK, body)
}
