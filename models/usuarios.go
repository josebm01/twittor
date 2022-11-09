package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	// omitempty, ignora datos vacios       bsoln: datos de entrada ---- json: datos de salida
	ID              primitive.ObjectID `bson: "_id, omitempty" json:"id"`
	Nombre          string             `bson: "nombre, omitempty" json:"nombre"`
	Apellidos       string             `bson: "apellidos, omitempty" json:"apellidos"`
	FechaNacimiento time.Time          `bson: "fechaNacimiento, omitempty" json:"fechaNacimiento"`
	Email           string             `bson: "email, json:"email"`
	Password        string             `bson: "password, json:"password, omitempty" ` // Devuelve el omitempery para no mostrar la contraseña
	Avatar          string             `bson: "avatar, json:"avatar, omitempty" `
	Banner          string             `bson: "banner, json:"biografia, omitempty" `
	Biografia       string             `bson: "biografia, json:"biografia, omitempty" `
	Ubicacion       string             `bson: "ubicacion, json:"ubicacion, omitempty" `
	SitioWeb        string             `bson: "sitioweb, json:"sitioweb, omitempty" `
}
