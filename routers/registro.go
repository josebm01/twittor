package routers

import (
	"encoding/json"
	"net/http"

	"github.com/josebm01/twittor/bd"
	"github.com/josebm01/twittor/models"
)

// Crear en la bd el registro de usuario
func Registro(w http.ResponseWriter, r *http.Request) {

	//objeto, modelo de usuario
	var t models.Usuario

	//cuerpo de body, es un objeto de tipo Stream que es un dato que solo se puede leer una vez y despues se destruye de la memoria
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe ser de al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.VerificarExistenciaUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "El email ya existe en otro usuario", 400)
		return
	}

	_, status, err := bd.InsertarRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro del usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logró insertar el usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
