package bd

import (
	"context"
	"time"

	"github.com/josebm01/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

func VerificarExistenciaUsuario(email string) (models.Usuario, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	// Devuelve un map string en formato bson
	condicion := bson.M{"email": email}

	var resultado models.Usuario

	// Búsqueda de un solo registro, para que devuelva uno solo y no una colección
	err := col.FindOne(ctx, condicion).Decode(&resultado) // Creando el json a través del puntero de la variable
	ID := resultado.ID.Hex()                              // Convirtiendo el objectId en string para devolverlo como tercer parámetro

	if err != nil {
		return resultado, false, ID
	}

	// Devolviendo el documento de mongo formateado en el modelo
	return resultado, true, ID

}
