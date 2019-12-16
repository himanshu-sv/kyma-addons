package handlers

import (
	"github.com/abbi-gaurav/kyma-connectors/components/github/event-gw/internal/incoming"
	"github.com/abbi-gaurav/kyma-connectors/components/github/event-gw/internal/logger"
	"github.com/abbi-gaurav/kyma-connectors/components/github/event-gw/internal/model/errors"
	"github.com/abbi-gaurav/kyma-connectors/components/github/event-gw/internal/outgoing"
	"net/http"
)

type EventPublisher struct {
	eventForwarder *outgoing.EventForwarder
	ip             *incoming.InboundProcessor
}

func NewEventPublisher() *EventPublisher {
	return &EventPublisher{
		ip:             incoming.New(),
		eventForwarder: outgoing.NewEventForwarder(),
	}
}

func (ep *EventPublisher) EventHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		kymaEvent, err := ep.ip.Process(r)
		if err != nil {
			errors.HandleError(w, err, errors.InternalError)
			return
		}
		logger.Logger.Infow("mapped payload", "kyma event type", kymaEvent.EventType)

		resp, err := ep.eventForwarder.Forward(kymaEvent)
		if err != nil {
			errors.HandleError(w, err, errors.InternalError)
			return
		}

		logger.Logger.Infow("Received response for event publishing", "response", resp)

		w.WriteHeader(http.StatusOK)
	})
}
