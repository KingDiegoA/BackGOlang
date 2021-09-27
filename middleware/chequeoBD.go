package middleware

import (
	"net/http"

	"github.com/KingDiegoA/BackGOlang/bd"
)

func CheckConn(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConn() == 0 {
			http.Error(w, "Conexion Fallida", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
