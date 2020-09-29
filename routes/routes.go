package routes

import (
	"crud/controllers"

	"github.com/gorilla/mux"
)

// SetRoutes agrega las rutas
func SetRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/api").Subrouter()
	// Contacts
	subRouter.HandleFunc("/contacts/{id}", controllers.GetContact).Methods("GET")
	subRouter.HandleFunc("/contacts", controllers.GetContacts).Methods("GET")
	subRouter.HandleFunc("/contacts", controllers.StoreContact).Methods("POST")
	subRouter.HandleFunc("/contacts/{id}", controllers.UpdateContact).Methods("PUT")
	subRouter.HandleFunc("/contacts/{id}", controllers.DeleteContact).Methods("DELETE")
	// Users
	subRouter.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	subRouter.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	subRouter.HandleFunc("/users", controllers.StoreUser).Methods("POST")
	subRouter.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	subRouter.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
}
