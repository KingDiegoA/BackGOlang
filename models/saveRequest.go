package models

import (
	"time"
)

/*Formato que tendra nuestra Solicitud en la BD*/
type SaveRequest struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
