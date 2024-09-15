package firebase

import (
	"context"
	"encoding/json"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"

	pubsub "github.com/tatangharyadi/pos-common/common/pubsub"
)

type QrPayment struct {
	Token     string `json:"token"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	ReceiptId string `json:"receipt_id"`
	Status    string `json:"status"`
}

func (h Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var body pubsub.PubSubMessage
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		h.Logger.Error().Msgf("error decoding request body: %v\n", err)
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	var qrPayment QrPayment
	json.Unmarshal(body.Message.Data, &qrPayment)

	h.Logger.Info().Msg("firebase: SendMessage sending...")

	config := &firebase.Config{ProjectID: h.Env.FirebaseProjectId}
	ops := []option.ClientOption{
		option.WithCredentialsJSON([]byte(h.Env.FirebaseServiceAccountKey)),
	}
	app, err := firebase.NewApp(context.Background(), config, ops...)
	if err != nil {
		h.Logger.Fatal().Msgf("error initializing app: %v\n", err)
	}

	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		h.Logger.Fatal().Msgf("error getting Messaging client: %v\n", err)
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: qrPayment.Title,
			Body:  qrPayment.Body,
		},
		Data: map[string]string{
			"orderId": "ID123",
			"status":  "SUCCESS",
		},
		Token: qrPayment.Token,
	}
	response, err := client.Send(ctx, message)
	if err != nil {
		h.Logger.Fatal().Msgf("error sending message: %v\n", err)
	}

	h.Logger.Info().Msgf("firebase: SendMessage success %v\n", response)
}
