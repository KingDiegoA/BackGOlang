package main

import (
	"log"

	"github.com/KingDiegoA/BackGOlang/bd"
	"github.com/KingDiegoA/BackGOlang/handlers"
)

func main() {
	if bd.CheckConn() == 0 {
		log.Fatal("Sin Conexion")
		return
	}
	handlers.Ejecutar()
}
