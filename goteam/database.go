package goteam

import (
"github.com/gin-gonic/gin"
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
"os"
"fmt"
)

// Consulta todos los registros de la base de datos
// token - Token de autorizacion
// query - Query de consulta
// obj - Estructura objeto de consulta
func Select(token string, query interface{}, obj interface{}) (int, interface{}) {
	if Auth(token){
		
		session := Connect();
		defer session.Close()
		collection := session.DB("my_bank_db").C("MensajesGo")
		
		err := collection.Find(query).All(&obj)
		if err != nil {
			panic(err)
		} else {
			return 200, obj
		}
	}
	return 404, bson.M{"error":  "permiso denegado"}
}

// Consulta un registro de la base de datos
// token - Token de autorizacion
// query - Query de consulta
// obj - Estructura objeto de consulta
func SelectOne(token string, query interface{}, obj interface{}) (int, interface{}) {
	if Auth(token){
		session := Connect();
		defer session.Close()
		collection := session.DB("my_bank_db").C("MensajesGo")
	
		err := collection.Find(query).One(&obj)
		if err != nil {
			return 404, gin.H{"auth":	"permiso concedido", "error":  "registro no encontrado"}
		} else {
			return 200, obj
		}
	}
	return 404, bson.M{"error":  "permiso denegado"}
}

// Inserta la base de datos
// token - Token de autorizacion
// query - Objeto a insertar
func Insert(token string, query interface{}) (int, interface{}) {
	if Auth(token){
		session := Connect();
		defer session.Close()
		collection := session.DB("my_bank_db").C("MensajesGo")
	
		err := collection.Insert(query)
		if err == nil {
			return 200, "Mensaje guardado!!!"
		}
	}
	return 404, bson.M{"error":  "permiso denegado"}
}

// Conecta a la base de datos
// Definir variable de entorno en Heroku: Settings, Config Vars
// Definir variable de entorno local: echo "export MONGO_URL=mongodb://goteam:goteam@ds019471.mlab.com:19471/my_bank_db" >> ~/.bashrc
func Connect() (session *mgo.Session) {
	connectURL := os.Getenv("MONGO_URL")
	session, err := mgo.Dial(connectURL)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	session.SetSafe(&mgo.Safe{})
	return session
}