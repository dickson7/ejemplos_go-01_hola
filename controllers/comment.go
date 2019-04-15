package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

// CommentGetAll se obtienen todos los comentarios
func CommentGetAll(w http.ResponseWriter, r *http.Request) {
	comments := []models.Comment{}
	m := models.Message{}
	user := models.User{}
	//vote := models.Vote{}

	r.Context().Value(&user)

	// /api/comments/?order=votes&idlimit=10
	// Con esta linea obtenemos lo que esta luego del / en la ruta
	vars := r.URL.Query()
	db := configuration.GetConnection()
	defer db.Close()
	// consulta de comentarios cComments
	cComment := db.Where("parent_id = 0")
	if order, ok := vars["order"]; ok {
		if order[0] == "votes" {
			cComment = cComment.Order("votes desc, created_at desc")
		}
	} else {
		if idlimit, ok := vars["idlimit"]; ok {
			registerByPage := 30
			offset, err := strconv.Atoi(idlimit[0])
			if err != nil {
				log.Println("Error:", err)
			}
			// TRAE TODOS LOS REGISTROS LIMITADOR EN 30 donde el id se encuentre entre 2 valores
			cComment = cComment.Where("id BETWEEN ? AND ?", offset-registerByPage, offset)
		}
		// venga o no venga un limite de registros se va a ordenar
		cComment = cComment.Order("id desc")
	}
	//ejecutamos la consulta y guardamos en un slice
	cComment.Find(&comments)
	// devolvemos todos los comentarios en un json
	j, err := json.Marshal(comments)
	if err != nil {
		m.Code = http.StatusInternalServerError
		m.Message = "Error al convertir los comentarios en json"
		commons.DisplayMessage(w, m)
		return
	}
	if len(comments) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		m.Code = http.StatusNoContent
		m.Message = "No se encontraron comentarios"
		commons.DisplayMessage(w, m)
	}
}
