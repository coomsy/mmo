package v1

import (
	"fmt"
	"net/http"

	"github.com/coomsy/mmo/pkg/log"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
}

// Default
func SetupHandler() *Handler {
	return &Handler{}
}

func (h Handler) GetUser(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
	if _, err := rw.Write([]byte("Login\n")); err != nil {
		log.Error("login error")
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func (h Handler) GetUserById(rw http.ResponseWriter, req *http.Request) {
	userId := chi.URLParam(req, "userId")

	rw.WriteHeader(http.StatusOK)

	payload := fmt.Sprintf("User id: %s", userId)
	if _, err := rw.Write([]byte(payload)); err != nil {
		log.Error("login error")
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
