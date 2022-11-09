package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/josebm01/twittor/middlewares"
	"github.com/josebm01/twittor/routers"
	"github.com/rs/cors"
)

// Se definen las rutas que se manejarán, seteo del puerto
func Manejadores() {
	// mux captura el HTTP y maneja la response y request para procesarla y enviar una respuesta al navegador
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.VerificarBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")

	// Definiendo un puerto por defecto, desde fuera se puede asignar otro
	if PORT == "" {
		PORT = "8080"
	}

	// Cors deja que  otras IP puedan acceder a la DB
	handler := cors.AllowAll().Handler(router)        // Permiso a todos
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) // El servidor escucha el puerto para ver los llamados de las peticiones
}
