package server

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/teramono/utilities/pkg/broker"
	"github.com/teramono/utilities/pkg/setup"
)

type BackendServer struct {
	setup.CommonSetup
}

func NewBackendServer(setup setup.CommonSetup) (BackendServer, error) {
	return BackendServer{
		CommonSetup: setup,
	}, nil
}

func (server *BackendServer) LogsVersion() uint {
	return server.Config.Broker.Subscriptions.Logs.Version
}

func (server *BackendServer) ActivateSubscriptions() error {
	// Create channels for subscribed subjects.
	runCh := make(chan *nats.Msg, server.BrokerClient.Opts.SubChanLen)
	runSub, err := server.BrokerClient.ChanQueueSubscribe(
		broker.GetWorkspacesSubjectByEngine(broker.EngineBackend, &server.Config, "run"),
		broker.GetWorkspacesResponderGroupByEngine(broker.EngineBackend, &server.Config, "run"),
		runCh,
	)
	if err != nil {
		return err
	}
	defer runSub.Unsubscribe()

	fmt.Println(">>> Subscription", broker.GetWorkspacesSubjectByEngine(broker.EngineBackend, &server.Config, "run"))
	fmt.Println(">>> Subscriptions set up!")

	// Listen to subscribed messages.
	for msg := range runCh {
		go broker.PanicWrap(msg, server.Run)
	}

	return nil
}
