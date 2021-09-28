package middleware

import (
	"net/http"

	"github.com/KingDiegoA/BackGOlang/routers"
)

/*ValidarJWT permite validar el JWT que nos viene en la petición */
func ValidarJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token ! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
