package models

import "github.com/jinzhu/gorm"

// Vote este modelos nos funciona como pivote para almacernar y saber quin voto que comentario y saber si el voto fue negativo o positivo
type Vote struct {
	gorm.Model
	CommentID uint `json:"commentId" gorm:"not null"`
	UserID    uint `json:"userId" gorm:"not null"`
	Value     bool `json:"value" gorm:"not null"`
}
