package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tatangharyadi/service/messaging/common/configs"
	"github.com/tatangharyadi/service/messaging/pkg/firebase"
)

func main() {
	env, logger := configs.InitEnv()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info().Msg("Hello World from service-messaging")
	})

	h := firebase.Handler{
		Env:    env,
		Logger: logger,
	}
	r.Mount("/firebase", h.Routes())

	logger.Info().Msgf("Listening %s mode:%s", env.AppEnv, env.AppPort)
	addr := ":" + env.AppPort
	http.ListenAndServe(addr, r)
}
