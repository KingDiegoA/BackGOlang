package routers

import (
	"encoding/json"
	"net/http"

	"github.com/KingDiegoA/BackGOlang/bd"
	"github.com/KingDiegoA/BackGOlang/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos"+err.Error(), 400)
		return
	}
	if len(t.Correo) == 0 {
		http.Error(w, "Debe ingresar el Email", 400)
		return
	}
	if len(t.Nacionalidad) == 0 {
		http.Error(w, "Debe ingresar la Nacionalidad", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password debe tener minimo 6 Caracteres", 400)
		return
	}

	_, encontrado, _ := bd.CheckUser(t.Correo)
	if encontrado == true {
		http.Error(w, "Usuario ya Existe", 400)
		return
	}
	_, status, _ := bd.InsertDats(t)
	if err != nil {
		http.Error(w, "Error al intentar realizar el Registro"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Registro Fallido", 400)
		return
	}
	if status == false {
		http.Error(w, "Registro Exitoso", 201)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
