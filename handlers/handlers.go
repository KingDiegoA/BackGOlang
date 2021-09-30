package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/KingDiegoA/BackGOlang/middleware"
	"github.com/KingDiegoA/BackGOlang/routers"
)

func Ejecutar() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middleware.CheckConn(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckConn(routers.Login)).Methods("POST")
	router.HandleFunc("/perfil", middleware.CheckConn(middleware.ValidarJWT(routers.Profile))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middleware.CheckConn(middleware.ValidarJWT(routers.UpdateProfile))).Methods("PUT")
	//Rutas de Solicitudes(Request)
	router.HandleFunc("/solicitudes", middleware.CheckConn(middleware.ValidarJWT(routers.SaveRequest))).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
