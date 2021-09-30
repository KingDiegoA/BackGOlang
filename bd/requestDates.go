package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/KingDiegoA/BackGOlang/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*InsertRequest graba el Request en la BD */
func InsertRequest(t models.SaveRequest) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("mongodb")
	col := db.Collection("solicitudes")

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}
	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
