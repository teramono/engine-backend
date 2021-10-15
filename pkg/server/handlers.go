package server

import (
	"fmt"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/teramono/utilities/pkg/broker"
	"github.com/teramono/utilities/pkg/logs"
	"github.com/teramono/utilities/pkg/messages"
	"github.com/teramono/utilities/pkg/request"
)

func (server *BackendServer) Run(msg *nats.Msg) {
	// Grab msg data.
	msgData, err := broker.NewMsgData(msg)
	if err != nil {
		server.sendServerErrorResponse(msg, messages.ProblemProcessingRequest)
		logs.Panicf(server, "%s: %v", messages.UnableToGetDataFromMessage, err)
		return
	}

	// Get workspace id from data header.
	workspaceIDs := msgData.Headers[request.WorkspaceIDHeader]
	if len(workspaceIDs) < 1 {
		server.sendServerErrorResponse(msg, messages.ProblemProcessingRequest)
		logs.Panicf(server, "%s: %v", messages.InvalidWorkspaceIDHeader.String(), err)
		return
	}
	workspaceID := workspaceIDs[0]

	// Get canonical workspace path.
	canonWorkspacePath, err := server.getWorkspaceCanonicalPath(workspaceID)
	if err != nil {
		server.sendServerErrorResponse(msg, messages.ProblemProcessingRequest)
		logs.Panicf(server, "%s: %v", messages.UnableToGetCanonicalWorkspacePath, err)
		return
	}

	// Get path suffix from url.
	pathFromURLSuffix := strings.TrimPrefix(msgData.URL.Path, "/r/")

	fmt.Println(">>> pathFromURLSuffix", pathFromURLSuffix)

	// Fetch surl manifest.
	manifest, err := server.fetchSurlManifest(canonWorkspacePath, pathFromURLSuffix)
	if err != nil {
		server.sendServerErrorResponse(msg, messages.ProblemProcessingRequest)
		logs.Panicf(server, "%s: %v", messages.UnableToFetchSurlManifest, err)
		return
	}

	fmt.Println(">>> manifest", manifest)

	// Run auth script.
	authSuccessful, err := server.runAuthScript(canonWorkspacePath)
	if err != nil {
		server.sendUserErrorResponse(msg, messages.ProblemProcessingRequest)
		logs.Panicf(server, "%s: %v", messages.UnableToRunAuthScriptSuccesfully, err)
		return
	}
	if !authSuccessful {
		server.sendUserErrorResponse(msg, messages.UnableToAuthenticate)
		logs.Panicf(server, "%s: %v", messages.UnableToAuthenticate, err)
		return
	}

	fmt.Println(">>> authSuccessful", authSuccessful)

	// Run middleware scripts.
	middlewareSuccessful, err := server.runMiddlewareScripts(canonWorkspacePath, &manifest)
	if err != nil {
		server.sendServerErrorResponse(msg, messages.ProblemProcessingRequest)
		logs.Panicf(server, "%s: %v", messages.UnableToRunMiddlewareScriptsSuccessfully, err)
		return
	}
	if !middlewareSuccessful {
		server.sendUserErrorResponse(msg, messages.ProblemProcessingRequest)
		logs.Panicf(server, "%s: %v", messages.UnableToRunIndexScriptSuccessfully, err)
		return
	}

	fmt.Println(">>> middlewareSuccessful", middlewareSuccessful)

	// Run surl index script.
	resp, err := server.runSurlIndexScript(canonWorkspacePath, pathFromURLSuffix)
	if err != nil {
		server.sendUserErrorResponse(msg, messages.ProblemProcessingRequest)
		logs.Panicf(server, "%s: %v", messages.UnableToAuthenticate, err)
		return
	}

	fmt.Println(">>> resp", string(resp))

	// Return a response.
	msg.Respond(resp)
}
