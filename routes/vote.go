package routes

import (
	"github.com/dickson7/proyectogocomentarios/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetVoteRouter es la ruta para el registro de un votp
func SetVoteRouter(router *mux.Router) {
	prefix := "/api/votes"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.VoteRegister).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.HandlerFunc(controllers.ValidateToken),
			negroni.Wrap(subRouter),
		),
	)
}
