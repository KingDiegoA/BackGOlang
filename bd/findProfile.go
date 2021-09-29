package bd

import (
	"context"
	"time"

	"github.com/KingDiegoA/BackGOlang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*FindProfile busca un profile en la BD */
func FindProfile(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoC.Database("mongodb")
	col := db.Collection("usuarios")

	var profile models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&profile)
	profile.Password = ""
	if err != nil {
		return profile, err
	}
	return profile, nil
}
