package main

import (
"github.com/zabawaba99/firego"
"github.com/gin-gonic/gin"
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
"os"
"strconv"
)

// Direccion de la app en Firebase (definir variable de entorno en Heroku)
var APP_URL = os.Getenv("APP_URL")

// Estructura Producto
type Producto struct {
	Nombre string
	Tipo   string
	Saldo  int
}

// Estructura Cliente
type Cliente struct {
	Id              bson.ObjectId `bson:"_id" json:"Id"`
	NombreCompleto  string `bson:"NombreCompleto" json:"NombreCompleto"`
	TipoDocumento   string `bson:"tipoDocumento" json:"TipoDocumento"`
	NumeroDocumento int `bson:"NumeroDocumento" json:"NumeroDocumento"`
	EjecutivoACargo string `bson:"ejecutivoAcargo" json:"EjecutivoACargo"`
	Productos       []Producto `bson:"productos" json:"Productos"`
}

// Define las rutas de la API y la ejecuta
func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("/clientes/:token", GetClientes)
		v1.GET("/cliente/:documento/:token", GetCliente)
		// v1.POST("/usuarios", PostUser)
		// v1.PUT("/usuarios/:id", UpdateUser)
		// v1.DELETE("/usuarios/:id", DeleteUser)
	}

	r.Run()
}

// Consulta la base de datos y retorna toda la coleccion de clientes
func GetClientes(ginContext *gin.Context) {
	session := connect();
	defer session.Close()
	collection := session.DB("my_bank_db").C("Clientes")
	token := ginContext.Params.ByName("token")

	if auth(token){
		clientes := []Cliente{}
		err := collection.Find(nil).All(&clientes)
		if err != nil {
			panic(err)
		}
		ginContext.JSON(200, clientes)
	}else{
		ginContext.JSON(404, gin.H{
			"error":  "permiso denegado",
			})
	}
}

// Consulta un documento de cliente en la base de datos y si existe retorna su conjunto de datos
func GetCliente(ginContext *gin.Context) {
	numeroDocumento := ginContext.Params.ByName("documento")
	numero, _ := strconv.ParseInt(numeroDocumento, 0, 64)
	token := ginContext.Params.ByName("token")

	if auth(token){
		session := connect();
		defer session.Close()
		collection := session.DB("my_bank_db").C("Clientes")

		cliente := Cliente{}
		err := collection.Find(bson.M{"NumeroDocumento": numero}).One(&cliente)
		if err != nil {
			ginContext.JSON(404, gin.H{
				"auth":	"permiso concedido",
				"error":  "registro no encontrado",
				})
		}else{
			ginContext.JSON(200, cliente)
		}

	}else {
		ginContext.JSON(404, gin.H{
			"error":  "permiso denegado",
			})
	}
}

// Conecta a la base de datos
func connect() (session *mgo.Session) {
	connectURL := "mongodb://goteam:goteam@ds019471.mlab.com:19471/my_bank_db"
	session, err := mgo.Dial(connectURL)
	if err != nil {
		// fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	session.SetSafe(&mgo.Safe{})
	return session
}

// Autentica un token de Google en Firebase
// Argumentos:
// token: token generado por Google
func auth(token string) bool {
	f := firego.New(APP_URL, nil)
	f.Auth(token)
	return processResponse(f)
}

// Procesa la respuesta de una peticion a Firebase
// Argumentos:
// f: instancia de Firebase con que se hizo peticion
func processResponse(f *firego.Firebase) bool {
	var v map[string]interface{}
	if err := f.Value(&v); err != nil {
		return false
	}
	return true
}
