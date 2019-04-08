package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/dickson7/proyectogocomentarios/migration"
	"github.com/dickson7/proyectogocomentarios/routes"
)

// para ejecutar la aplicacion y ejecutar la migracion
// ./proyectogocomentarios --migrate yes
// proyectogocomentarios.exe --migrate yes
// de esta manera se generan las tablas en la BD
// si no ejecuta migrate el valor seria no y no se llama a migrate

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzo la migración...")
		migration.Migrate()
		log.Println("Finalizo la migración.")
	}

	// inicializacion de la rutas
	router := routes.InitRoutes()

	// inicializa los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// Inicia el servidor
	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}
	log.Println("Iniciado el servidor en http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecucion del programa")

}
