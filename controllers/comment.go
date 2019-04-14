package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dickson7/proyectogocomentarios/commons"
	"github.com/dickson7/proyectogocomentarios/configuration"
	"github.com/dickson7/proyectogocomentarios/models"
)

// CommentCreate permite registrar comentarios en la base de datos
func CommentCreate(w http.ResponseWriter, r *http.Request) {
	// se valida el comentario
	comment := models.Comment{}
	m := models.Message{}

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al leer el comentario: %s", err)
		commons.DisplayMessage(w, m)
		return
	}
	// si el comentario fue leido con exito y no dio ningun error
	// nos conectamos a la base de datos y creamos el registro

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&comment).Error
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al registrar el comentario: %s", err)
		commons.DisplayMessage(w, m)
		return
	}

	m.Code = http.StatusCreated
	m.Message = "Comentario creado con éxito"
	commons.DisplayMessage(w, m)
}
