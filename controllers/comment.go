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

	// traemos el usuario del contexto
	user := models.User{}
	user, _ = r.Context().Value("user").(models.User)

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		m.Code = http.StatusBadRequest
		m.Message = fmt.Sprintf("Error al leer el comentario: %s", err)
		commons.DisplayMessage(w, m)
		return
	}
	comment.UserID = user.ID

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
	vote := models.Vote{}

	user, _ = r.Context().Value("user").(models.User)

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
	// recorremos el slice para agregar el usuario a cada comentario
	for i := range comments {
		db.Model(&comments[i]).Related(&comments[i].User)
		// en este recorrido tambien trae el password entonces debemos enviarlo vacio
		comments[i].User[0].Password = ""
		// tambien aprovechamos para buscar los comentarios +hijos
		comments[i].Children = commentGetChilden(comments[i].ID)

		// se busca le voto del usuario en sesion
		vote.CommentID = comments[i].ID
		vote.UserID = user.ID
		count := db.Where(&vote).Find(&vote).RowsAffected
		if count > 0 {
			if vote.Value {
				comments[i].HasVote = 1
			} else {
				comments[i].HasVote = -1
			}
		}

	}

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

// commentGetChilden comentarios hijos funcion interna
func commentGetChilden(id uint) (children []models.Comment) {
	db := configuration.GetConnection()
	defer db.Close()

	db.Where("parent_id = ?", id).Find(&children)
	//buscamos el usuario que hizo el comentario hijo
	for i := range children {
		db.Model(&children[i]).Related(&children[i].User)
		// en este recorrido tambien trae el password entonces debemos enviarlo vacio
		children[i].User[0].Password = ""
	}
	return
}
