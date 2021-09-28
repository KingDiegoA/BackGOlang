package jwt

import (
	"time"

	"github.com/KingDiegoA/BackGOlang/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GeneroJWT genera el encriptado con JWT */
func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte("BackGolangPass")

	payload := jwt.MapClaims{
		"correo":       t.Correo,
		"nombre":       t.Nombre,
		"edad":         t.Edad,
		"nacionalidad": t.Nacionalidad,
		"fingreso":     t.Fingreso,
		"ftermino":     t.Ftermino,
		"_id":          t.ID.Hex(),
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
