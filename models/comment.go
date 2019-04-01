package models

import "github.com/jinzhu/gorm"

// Comment contiene los comentarios realizados por el usuario
type Comment struct {
	gorm.Model
	UserID   uint   `json:"userId"`
	ParentID uint   `json:"parentId"`
	Votes    int32  `json:"votes"`
	Content  string `json:"content"`
	// hasvote verifica que un usuario no vote mas de una ves por un comentarios
	HasVote int8 `json:"hasVote" gorm:"-"`
	// traemos un slice de usuario pero en realizada solo trae un usuario que seria el que creo el comentario
	User []User `json:"user,omitempty"`
	// un comentario puede tener varios comentarios hijos
	Children []Comment `json:"children,omitempty"`
}
