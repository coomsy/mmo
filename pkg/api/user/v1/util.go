package v1

import (
	"net/http"

	"github.com/coomsy/mmo/pkg/log"
)

type Handler struct {
}

func SetupHandler() *Handler {
	return &Handler{}
}

func (h Handler) Login(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)

	if _, err := rw.Write([]byte("Login\n")); err != nil {
		log.Error("login error")
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
