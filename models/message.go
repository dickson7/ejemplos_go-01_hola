package models

//Message contiene los mensajes de alerta para la api
type Message struct {
	Message string `jeson:"message"`
	Code    int    `json:"code"`
}
