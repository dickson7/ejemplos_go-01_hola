package main

import (
	"flag"
	"log"

	"github.com/dickson7/proyectogocomentarios/migration"
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
}
