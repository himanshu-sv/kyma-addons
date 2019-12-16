package incoming

import (
	"github.com/abbi-gaurav/kyma-connectors/components/github/event-gw/internal/config"
	"github.com/abbi-gaurav/kyma-connectors/components/github/event-gw/internal/model/events"
	"github.com/google/go-github/github"
	"net/http"
)

type InboundProcessor struct {
	secretKey []byte
}

func New() *InboundProcessor {
	return &InboundProcessor{
		secretKey: []byte(*config.GlobalConfig.SecretKey),
	}
}

func (ip *InboundProcessor) Process(r *http.Request) (*events.KymaEvent, error) {
	payload, err := github.ValidatePayload(r, ip.secretKey)
	if err != nil {
		return nil, err
	}
	eventType := github.WebHookType(r)

	return events.Map(payload, eventType)

}
