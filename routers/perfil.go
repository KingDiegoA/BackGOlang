package routers

import (
	"encoding/json"
	"net/http"

	"github.com/KingDiegoA/BackGOlang/bd"
)

/*Profile permite extraer los valores del Profile */
func Profile(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el ID", http.StatusBadRequest)
		return
	}

	profile, err := bd.FindProfile(ID)
	if err != nil {
		http.Error(w, "Error al intentar buscar el registro "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
