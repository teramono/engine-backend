package engine

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login ...
func (server *BackendServer) Login(c *gin.Context) {
}

// Run ...
func (server *BackendServer) Run(c *gin.Context) {
	// TODO: Will eventually use grpc.
	var body map[string]interface{}
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
	}

	// TODO: Contruct request.

	c.JSON(http.StatusBadRequest, body)
}
