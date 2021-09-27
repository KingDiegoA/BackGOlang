package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoC es el objeto de conexión a la BD */
var MongoC = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:1234@cluster0.ufxyn.mongodb.net/test")

/*ConectarBD es la función que me permite conectar la BD */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Print("Conexion Exitosa...")
	return client
}

/*CheckConn es el Ping a la BD */
func CheckConn() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
