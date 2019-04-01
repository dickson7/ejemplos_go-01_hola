package models

import "github.com/jinzhu/gorm"

// User contiene informacion de usuario
type User struct {
	gorm.Model
	// como se va a consumi via json no se utiliza mayuscula
	// y el orm le decimos que no puede ser nulo y que va a ser unico
	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Fullname string `json:"fullname" gorm:"not null"`
	// el json le vamos a decir que si el password esta vacio no lo envie y que no lo envie si esta vacio
	Password string `json:"password,omitempty" gorm:"not null;type:varchar(255)"`
	// la confirmacion del password no es necesario guardarla en la BD entonces le decimos al gorm que la omita
	ConfirmPassword string    `json:"confirmPassword,omitempty" gorm:"-"`
	Picture         string    `json:"picture"`
	Comments        []Comment `json:"commets,omitemty"`
}
