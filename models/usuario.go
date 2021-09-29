package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre       string             `bson:"nombre" json:"nombre,omitempty"`
	Edad         int                `bson:"edad" json:"edad,omitempty"`
	Correo       string             `bson:"correo" json:"correo"`
	Nacionalidad string             `bson:"nacionalidad" json:"nacionalidad,omitempty"`
	Empresa      string             `bson:"empresa" json:"empresa,omitempty"`
	Fingreso     time.Time          `bson:"fingreso" json:"fingreso,omitempty"`
	Ftermino     time.Time          `bson:"ftermino" json:"ftermino,omitempty"`
	Password     string             `bson:"password" json:"password,omitempty"`
	Avatar       string             `bson:"avatar" json:"avatar,omitempty"`
	Banner       string             `bson:"banner" json:"banner,omitempty"`
}
