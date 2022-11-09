package middlewares

import (
	"net/http"

	"github.com/josebm01/twittor/bd"
)

// Devuelve una funcion
func VerificarBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.VerificarConexion() == 0 {
			// Recibe la respuesta y manda el mensaje con un status 500
			http.Error(w, "Conexión perdida con la base de datos", 500)
			// Detiene todo si da error
			return
		}

		// Si no da error pasa los valores recibidos
		next.ServeHTTP(w, r)
	}
}
