package bd

import (
	"context"
	"time"

	"github.com/KingDiegoA/BackGOlang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*UpdateProfile permite modificar el perfil del usuario */
func UpdateProfile(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("mongodb")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if int(u.Edad) > 0 {
		registro["edad"] = u.Edad
	}
	if len(u.Nacionalidad) > 0 {
		registro["nacionalidad"] = u.Nacionalidad
	}
	if len(u.Empresa) > 0 {
		registro["empresa"] = u.Empresa
	}
	registro["fingreso"] = u.Fingreso
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	registro["ftermino"] = u.Ftermino
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}

	updtString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}
