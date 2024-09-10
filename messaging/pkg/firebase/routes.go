package firebase

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/tatangharyadi/service/messaging/common/configs"
)

type Handler struct {
	Env    *configs.Env
	Logger zerolog.Logger
}

func (h Handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/sendmessage", h.SendMessage)
	return r
}
