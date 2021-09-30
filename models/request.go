package models

/*Request captura el Body, mensaje que nos llega */
type Request struct {
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
}
