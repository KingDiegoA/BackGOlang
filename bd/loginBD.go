package bd

import (
	"github.com/KingDiegoA/BackGOlang/models"
	"golang.org/x/crypto/bcrypt"
)

/*Login realiza el chequeo de login a la BD */
func Login(correo string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := CheckUser(correo)
	if encontrado == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
