package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KingDiegoA/BackGOlang/bd"
	"github.com/KingDiegoA/BackGOlang/jwt"
	"github.com/KingDiegoA/BackGOlang/models"
)

/*Registro realiza el registro en la BD */
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

/*Login realiza el login en la BD */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña inválidos "+err.Error(), 400)
		return
	}
	if len(t.Correo) == 0 {
		http.Error(w, "El email del usuario es requerido ", 400)
		return
	}
	documento, existe := bd.Login(t.Correo, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o Contraseña inválidos ", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar general el Token correspondiente "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}

/*Profile permite extraer los valores del perfil */
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

/*UpdateProfile modifica el perfil de usuario */
func UpdateProfile(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.UpdateProfile(t, IDUsuario)
	if err != nil {
		http.Error(w, "Error al modificar los datos. Reintente nuevamente "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar los datos del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
