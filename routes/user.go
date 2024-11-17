package routes

import (
	"elasticsearch/feature/user"
	"net/http"
)

// User mengatur rute terkait user
func User(mux *http.ServeMux, handler user.Handler) {
	mux.HandleFunc("/users", handler.CreateUser)
	mux.HandleFunc("/users/details", handler.GetUserDetails)
	mux.HandleFunc("/users/search", handler.SearchUsers)
	mux.HandleFunc("/users/delete", handler.DeleteUser)
	mux.HandleFunc("/users/update", handler.UpdateUser)
}
