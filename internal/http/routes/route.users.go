package routes

import "net/http"

type usersRoutes struct {
	router         *http.ServeMux
	authMiddleware http.Handler
}

// Meh
func SetupUsersRoute(router *http.ServeMux) {
	
}

func (r* usersRoutes) Routes() {

	r.router.HandleFunc("/v1/auth/login", )
}