package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
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

/*ViewRequest vista de las solicitudes */
func ViewRequest(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)
	respuesta, correcto := bd.ResponseRequest(ID, pag)
	if correcto == false {
		http.Error(w, "Error al leer las Solicitudes", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}

/*DeleteRequest permite borrar un Request determinado */
func DeleteRequest(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	err := bd.DeleteRequest(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar eliminar "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
