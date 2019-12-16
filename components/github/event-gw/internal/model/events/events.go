package events

import (
	"encoding/json"
	"github.com/abbi-gaurav/kyma-connectors/components/github/event-gw/internal/config"
	"github.com/gofrs/uuid"
	"time"
)

type Data map[string]interface{}

type KymaEvent struct {
	SourceID         *string `json:"source-id"`
	EventType        string  `json:"event-type"`
	EventTypeVersion string  `json:"event-type-version"`
	EventID          string  `json:"event-id"`
	EventTime        string  `json:"event-time"`
	Data             *Data   `json:"data"`
}

func Map(payload []byte, eventType string) (*KymaEvent, error) {
	eventId, _ := generateEventID()
	data, err := to(payload)
	if err != nil {
		return nil, err
	}
	return &KymaEvent{
		SourceID:         config.GlobalConfig.AppName,
		EventTypeVersion: "v1",
		EventTime:        time.Now().Format(time.RFC3339),
		Data:             data,
		EventID:          eventId,
		EventType:        eventType,
	}, nil
}

func generateEventID() (string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}

func to(payload []byte) (*Data, error) {
	var data Data
	err := json.Unmarshal(payload, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
