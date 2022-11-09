package main

import (
	"log"

	"github.com/josebm01/twittor/bd"
	"github.com/josebm01/twittor/handlers"
)

func main() {
	if bd.VerificarConexion() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}

	handlers.Manejadores()
}
