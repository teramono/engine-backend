package engine

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RunScriptURL ...
func (server *BackendServer) RunScriptURL(c *gin.Context) {
	// TODO: Will eventually use grpc.
	var body map[string]interface{}
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
	}

	// TODO: Contruct request.

	c.JSON(http.StatusBadRequest, body)
}
