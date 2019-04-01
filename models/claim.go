package models

// se le da un alias al paquete para no utilizar todo el nombre
import jwt "github.com/dgrijalva/jwt-go"

// Claim es una solicitud
type Claim struct {
	// no se usa el orm porque esta info no se va almacer en la BD
	// este Clain sirve para verificar si el usuario se autentico o no
	User `json:"user"`
	jwt.StandardClaims
}
