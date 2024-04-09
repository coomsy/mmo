package routes

import (
	"fmt"
	"net/http"
)

func Root(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		if _, err := rw.Write([]byte("pong\n")); err != nil {
			fmt.Println("error when writing response for /ping request")
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}
}
