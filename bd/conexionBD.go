package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Variable que se exportará a todos los archivos
var MongoCN = ConectarBD()

// Setear la URL de la conexión
var clientOptions = options.Client().ApplyURI("mongodb+srv://jose:jose@twittor.ibctqr8.mongodb.net/?retryWrites=true&w=majority")

// De vuelve un objeto de la conexion
func ConectarBD() *mongo.Client {
	// Conexion
	// context: Espacio de memoria para compartir informacion de la conexion a la bd a través de la ejecución y permite setear información
	// -TODO significa que no hay restricciones en la conexion
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		// .Error() muestra el mensaje del error en un string
		log.Fatal(err.Error())
		return client // Devolver el objeto aunque se encuentre vacio
	}

	// Para verificar si la base de datos se encuentra disponible
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión exitosa con la DB")
	return client

}

// Ping a la DB
func VerificarConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
