package firebase

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

func (h Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
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

	registrationToken := "diZjYadBT0KKj1YRZQwQp7:APA91bEnmxuYzvCdvx8uqIr0AlFbHfhPOBAQzdvVmi3JUfzzGaKeo56JIiyMTKmYtJOfx-2KiGNKM0OhBAjrwusFzGlWb7QhDwOGrKGOdgA3n5zycjjUzAeojqamw8NYxwzWp0NUHOyE"
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "FCM Message",
			Body:  "This is an FCM Message",
		},
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}
	response, err := client.Send(ctx, message)
	if err != nil {
		h.Logger.Fatal().Msgf("error sending message: %v\n", err)
	}

	h.Logger.Info().Msgf("firebase: SendMessage success %v\n", response)
}
