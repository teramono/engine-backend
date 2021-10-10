package server

import (
	"net/http"

	"github.com/nats-io/nats.go"
	"github.com/teramono/utilities/pkg/broker"
	"github.com/teramono/utilities/pkg/logs"
	"github.com/teramono/utilities/pkg/messages"
	"github.com/teramono/utilities/pkg/response"
)

func (server *BackendServer) sendErrorResponse(msg *nats.Msg, errMsg messages.ErrorMessage, statusCode int) {
	data, err := response.CreateErrorResponseBodyBytes(errMsg)
	if err != nil {
		logs.Panicf(server, "%s: %v", messages.UnableToConstructServerErrorResponse, err)
	}

	msgData, err := broker.JsonFromMsgData(
		broker.URL{},
		broker.Header{
			"Content-Type": {"application/json", "charset=utf-8"},
		},
		string(data),
		statusCode,
	)
	if err != nil {
		logs.Panicf(server, "%s: %v", messages.UnableToConstructServerErrorResponse, err)
	}

	if err = msg.Respond(msgData); err != nil {
		logs.Panicf(server, "%s: %v", messages.UnableToSendServerErrorResponse, err)
	}
}

func (server *BackendServer) sendServerErrorResponse(msg *nats.Msg, errMsg messages.ErrorMessage) {
	server.sendErrorResponse(msg, errMsg, http.StatusInternalServerError)
}

func (server *BackendServer) sendUserErrorResponse(msg *nats.Msg, errMsg messages.ErrorMessage) {
	server.sendErrorResponse(msg, errMsg, http.StatusBadRequest)
}
