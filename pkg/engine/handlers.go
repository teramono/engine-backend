package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teramono/utilities/pkg/messages"
	"github.com/teramono/utilities/pkg/request"
	"github.com/teramono/utilities/pkg/response"
)

// Login ...
func (server *BackendServer) Login(ctx *gin.Context) {
}

// Run ...
func (server *BackendServer) Run(ctx *gin.Context) {
	// Construct request.
	req, err := request.NewRequestFromContext(ctx)
	if err != nil {
		response.SetErrorResponse(
			ctx,
			messages.ErrorMessage(err.Error()), // TODO: User friendly message
		)
		return
	}

	// Fetch and run auth script.
	if err := server.fetchAndRunAuthScript(&req); err != nil {
		response.SetErrorResponse(
			ctx,
			messages.ErrorMessage(err.Error()), // TODO: User friendly message
		)
		return
	}

	// Fetch and bundle index script with middlewares using manifest
	// TODO: Run in goroutine. Send the response back via socket or callback URL.
	_, err = server.fetchAndRunBundledScript(&req)
	if err != nil {
		response.SetErrorResponse(
			ctx,
			messages.ErrorMessage(err.Error()), // TODO: User friendly message
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ //
		"message": "Script running...",
	})
}
