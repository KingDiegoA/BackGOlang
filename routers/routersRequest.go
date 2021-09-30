package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KingDiegoA/BackGOlang/bd"
	"github.com/KingDiegoA/BackGOlang/models"
)

/*SaveRequest permite grabar la solicitud en la base de datos */
func SaveRequest(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Request
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.SaveRequest{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertRequest(registro)
	if err != nil {
		http.Error(w, "Error al intentar insertar el registro"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar la Solicitud", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
