package engine

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teramono/utilities/pkg/database/models"
)

// Run ...
// TODO: Remove!
func (engine *BackendEngine) Run(c *gin.Context) {
	// TODO: Will eventually use grpc.
	var body map[string]interface{}
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
	}

	// ...
	_, err := strconv.Atoi(c.Query("workspace_id"))
	if err != nil {
		body["message"] = "invalid id query"
		c.JSON(http.StatusBadRequest, body)
		return
	}

	// ...
	if script, _ := body["script"].(string); script != "" {
		if err := engine.runScript(script); err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, body)
		return
	}

	// ...
	if scriptFile, _ := body["scriptFile"].(string); scriptFile != "" {
		if err := engine.runScriptFile(scriptFile); err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, body)
		return
	}

	c.JSON(http.StatusBadRequest, body)
}

// LogIn ...
func (engine *BackendEngine) LogIn(c *gin.Context) {
	// TODO: Will eventually use grpc.
	var body map[string]interface{}
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
	}

	// ...
	workspaceName := c.Query("workspace_name")
	fmt.Println("workspace name", workspaceName)
	if workspaceName == "" {
		body["message"] = "invalid name query"
		c.JSON(http.StatusBadRequest, body)
		return
	}

	// ...
	workspace := models.Workspace{Name: workspaceName}
	workspace, err := workspace.GetByName(&engine.setup.WorkspacesDB)
	if err != nil {
		body["message"] = err
		c.JSON(http.StatusBadRequest, body)
		return
	}

	fmt.Println("workspace id", workspace.ID)

	c.JSON(http.StatusOK, body)
}
