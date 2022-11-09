package bd

import "golang.org/x/crypto/bcrypt"

func EncriptarPassword(pass string) (string, error) {
	// Cantidad de 2 elevado al costo, para realizar la encriptación
	costo := 8
	// La función solo aceptas slices de bytes
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	// Procesa un slice de bytes y devuelve lo mismo
	return string(bytes), err
}
