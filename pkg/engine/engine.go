package engine

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BackendEngine struct {
	isolates []Isolate
}

func NewBackendEngine() (BackendEngine, error) {
	isolate, err := NewIsolate()
	if err != nil {
		return BackendEngine{}, err
	}

	return BackendEngine{
		isolates: []Isolate{isolate},
	}, nil
}

func (engine *BackendEngine) Listen() error {
	// TODO:
	//  raw - Execute functions sent in requests.
	//  fs - Execute functions linked in requests.
	//  middleware -  middleware functions are synchronous functions, their result must be gotten before proceeding.

	router := gin.Default()
	router.POST("/workspace/:id", func(c *gin.Context) {
		var body map[string]string

		if err := c.BindJSON(&body); err != nil {
			fmt.Println(err)
		}

		userScript := body["script"]

		workspaceID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println(err, workspaceID)
		}

		if err := engine.RunScript(userScript); err != nil {
			fmt.Println(err)
		}

		// TODO: Determine isolate to run or create new one.

		c.JSON(http.StatusOK, gin.H{
			"code":        userScript,
			"workspaceId": workspaceID,
		})
	})

	return router.Run()
}

func (engine *BackendEngine) GetIsolate() (Isolate, error) {
	// TODO: Must be a way for creating isolates.
	return engine.isolates[0], nil
}

func (engine *BackendEngine) RunMiddlewareScripts() error {
	// TODO:
	return nil
}

func (engine *BackendEngine) RunScript(script string) error {
	// TODO: Run synchronous middleware first.

	isolate, err := engine.GetIsolate()
	if err != nil {
		return err
	}

	// TODO:
	//  Run script in a separate goroutine. Research if multiple goroutines accessing an isolate memory is a bad thing./

	value, err := isolate.RunScript(script)
	if err != nil {
		return err
	}

	// TODO:
	//  Return result value as an object websocket message.

	fmt.Println("Value", value.String())

	return nil
}
