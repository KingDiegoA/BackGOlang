package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Formato que tendra nuestra Solicitud en la BD*/
type SaveRequest struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}

/*Request captura el Body, mensaje que nos llega */
type Request struct {
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
}

/*Response es la estructura con la que devolveremos las solicitudes */
type Response struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userid" json:"userId,omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
