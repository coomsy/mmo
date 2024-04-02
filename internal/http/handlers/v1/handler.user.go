package v1

import (
	"net/http"

	"github.com/coomsy/mmo/pkg/logger"
)

type userHandler struct {
}

func (h userHandler) Login(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		rw.WriteHeader(http.StatusOK)
		if _, err := rw.Write([]byte("Login\n")); err != nil {
			logger.Error("login error")
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}
}
