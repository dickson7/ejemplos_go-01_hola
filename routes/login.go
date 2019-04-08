package routes

import (
	"github.com/dickson7/proyectogocomentarios/controllers"
	"github.com/gorilla/mux"
)

// SetLoginRouter ruta para el login
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
