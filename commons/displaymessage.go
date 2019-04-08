package commons

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dickson7/proyectogocomentarios/models"
)

// DisplayMessage devuelve un mensa al cliente
func DisplayMessage(w http.ResponseWriter, m models.Message) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error al convertir el mensaje: %s", err)
	}
	w.WriteHeader(m.Code)
	w.Write(j)
}
