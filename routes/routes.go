package routes

import (
	"github.com/gorilla/mux"
)

// InitRoutes archivo principal de rutas que la inicializa
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetLoginRouter(router)
	SetUserRouter(router)
	SetCommentRouter(router)

	return router
}
