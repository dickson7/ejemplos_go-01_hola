package controllers

import (
	"context"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dickson7/proyectogocomentarios/commons"
	"github.com/dickson7/proyectogocomentarios/models"
)

//ValidateToken valida el token del cliente
func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var m models.Message
	token, err := request.ParseFromRequestWithClaims(
		r,
		request.OAuth2Extractor,
		&models.Claim{},
		func(t *jwt.Token) (interface{}, error) {
			return commons.PublicKey, nil
		},
	)
	if err != nil {
		m.Code = http.StatusUnauthorized
		switch err.(type) {
		case *jwt.ValidationError:
			vError := err.(*jwt.ValidationError)
			switch vError.Errors {
			case jwt.ValidationErrorExpired:
				m.Message = "Su token ha expirado"
				commons.DisplayMessage(w, m)
				return
			case jwt.ValidationErrorSignatureInvalid:
				m.Message = "La firma del token no coincide"
				commons.DisplayMessage(w, m)
				return
			default:
				m.Message = "Su token no es valido"
				commons.DisplayMessage(w, m)
				return
			}
		}
	}
	if token.Valid {
		// aprovechamos que ya tenemos el token extraido y extremos la informacion del usuario
		// utilizando el context
		ctx := context.WithValue(r.Context(), "user", token.Claims.(*models.Claim).User)
		// enviamos el request pero el request lo enviamos con un contexo que es la info del usuario
		next(w, r.WithContext(ctx))
	} else {
		m.Code = http.StatusUnauthorized
		m.Message = "Su token no es valido"
		commons.DisplayMessage(w, m)
	}
}
