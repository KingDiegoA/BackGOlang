package bd

import (
	"context"
	"time"

	"github.com/KingDiegoA/BackGOlang/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckUser(correo string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("mongodb")
	col := db.Collection("usuarios")

	condicion := bson.M{"correo": correo}
	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
