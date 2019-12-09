// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"context"
	"encoding/json"

	"github.com/getsentry/sentry-go"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// SendToSentry consumes a Pub/Sub message.
func SendToSentry(ctx context.Context, m PubSubMessage) error {
	sentry.Init(sentry.ClientOptions{
		Dsn: "INSERT_DSN_HERE",
	})

	// Parse JSON data into map
	var result map[string]interface{}
	json.Unmarshal(m.Data, &result)

	resource := result["resource"].(map[string]interface{})

	labels := resource["labels"].(map[string]interface{})

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		for k, v := range labels {
			if v.(string) != "" {
				scope.SetTag(k, v.(string))
			}
		}

		if result["severity"] == "CRITICAL" {
			scope.SetLevel(sentry.LevelFatal)
		} else {
			scope.SetLevel(sentry.LevelError)
		}

		if jsonPayload, ok := result["jsonPayload"]; ok {
			jsonPayload := jsonPayload.(map[string]interface{})

			for k, v := range jsonPayload {
				if k != "message" {
					scope.SetTag(k, v.(string))
				}
			}

			if message, ok := jsonPayload["message"]; ok {
				sentry.CaptureMessage(message.(string))
			}

		} else if textPayload, ok := result["textPayload"]; ok {
			sentry.CaptureMessage(textPayload.(string))
		}
	})

	return nil
}
