package migration

// este paquete se utiliza solo una vez para crear las tablas en la BD

import (
	"github.com/dickson7/proyectogocomentarios/configuration"
	"github.com/dickson7/proyectogocomentarios/models"
)

//Migrate se conecta a la BD y permie crear las tablas
func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	//Se crean los modelos
	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})
	// vote debe tener una llave unica para que no se repita el voto (se colocan los campos que haces la llave)
	db.Model(&models.Vote{}).AddUniqueIndex("comment_id_user_id_unique", "comment_id", "user_id")
}
