package bd

import (
	"context"
	"time"

	"github.com/josebm01/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Devuelve 3 valores
func InsertarRegistro(u models.Usuario) (string, bool, error) {

	// WithTimeout: Para que no supere un tiempo especificado,
	// background: lo que tra en la base de datos

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	// El defer se define al principio pero se ejecuta al final de la instrucción
	defer cancel() // al objeto devuelvo cancela el timeout (contexto) para no llenar la memoria

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	u.Password, _ = EncriptarPassword(u.Password)

	// Operacion en la base de datos, inserta un registro
	result, err := col.InsertOne(ctx, u)

	// Si ocurre un error se devuelve un id vacio, falso de que no fue satisfactoria y el error
	if err != nil {
		return "", false, err
	}

	// Obteniendo el objectID
	objID, _ := result.InsertedID.(primitive.ObjectID)

	// Convirtiendo de objectID a string
	return objID.String(), true, nil
}
