package v1

import (
	"encoding/json"
	"net/http"
	"time"

	guildV1 "github.com/coomsy/mmo/pkg/apis/guild/v1"
	"github.com/coomsy/mmo/pkg/log"

	"github.com/go-chi/chi/v5"
)

// Stuff for db conn etc
type Handler struct {
}

func SetupHandler() *Handler {
	return &Handler{}
}

func (h Handler) GetGuild(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	if _, err := rw.Write([]byte("Login\n")); err != nil {
		log.Error("login error")
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func (h Handler) GetGuildById(rw http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	rw.WriteHeader(http.StatusOK)

	guild, _ := json.Marshal(&guildV1.Guild{
		Id:        id,
		Name:      "2 ez",
		CreatedAt: time.Now(),
	})

	if _, err := rw.Write(guild); err != nil {
		log.Error("login error")
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
