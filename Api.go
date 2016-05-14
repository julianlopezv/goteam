package main

import (
"github.com/zabawaba99/firego"
"github.com/gin-gonic/gin"
"github.com/itsjamie/gin-cors"
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
"os"
"strconv"
"fmt"
"time"
)
//
// Direccion de la app en Firebase
// Definir variable de entorno en Heroku: Settings, Config Vars
// Definir variable de entorno local: echo "export APP_URL=https://radiant-inferno-2748.firebaseio.com" >> ~/.bashrc
var APP_URL = os.Getenv("APP_URL")

// Estructura Producto
type Producto struct {
	Nombre string `bson:"nombre" json:"tipo"`
	Tipo   string `bson:"tipo" json:"tipo"`
	Saldo  int `bson:"saldo" json:"saldo"`
}

// Estructura Cliente
type Cliente struct {
	Id              bson.ObjectId `bson:"_id" json:"_id"`
	NombreCompleto  string `bson:"nombre_completo" json:"nombre_completo"`
	TipoDocumento   string `bson:"tipo_doc" json:"tipo_doc"`
	NumeroDocumento int `bson:"documento" json:"documento"`
	EjecutivoEncargado string `bson:"ejecutivo_encargado" json:"ejecutivo_encargado"`
	Correo string `bson:"correo" json:"correo"`
	Productos       []Producto `bson:"productos" json:"productos"`
}

// Estructura Ubicacion
type Ubicacion struct {
	Longitud float64 `bson:"longitud" json:"longitud"`
	Latitud float64 `bson:"latitud" json:"latitud"`
}

// Estructura Ejecutivo
type Ejecutivo struct {
	Id              bson.ObjectId `bson:"_id" json:"_id"`
	IdEjecutivo string    `bson:"id_ejecutivo" json:"id_ejecutivo"`
	Nombre  string `bson:"nombre" json:"nombre"`
	Foto  string `bson:"foto" json:"foto"`
	TipoDocumento   string `bson:"tipo_doc" json:"tipo_doc"`
	NumeroDocumento int `bson:"documento" json:"documento"`
	Correo string `bson:"correo" json:"correo"`
	Celular int `bson:"celular" json:"celular"`
	Ubicacion Ubicacion `bson:"ubicacion" json:"ubicacion"`
}

// Define las rutas de la API y la ejecuta
func main() {
	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
		}))
	v1 := r.Group("")
	{
		v1.GET("/", ImOk)
		v1.GET("/clientes/:token", GetClientes)
		// v1.GET("/cliente/documento/:documento/:token", GetCliente)
		v1.GET("/cliente/:correo/:token", GetClientePorCorreo)
		v1.GET("/ejecutivo/:correo/:token", GetEjecutivoPorCorreoCliente)
		// v1.POST("/usuarios", PostUser)
		// v1.PUT("/usuarios/:id", UpdateUser)
		// v1.DELETE("/usuarios/:id", DeleteUser)
	}

	/*para correr en un puerto local*/
	//r.Run(":1337")
	r.Run()
}

// Meotodo que permite responder peticiones que vengan desde otros dominios


func ImOk(ginContext *gin.Context) {
	ginContext.JSON(200, gin.H{
		"status":  "Im OK!!!",
		})
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
		err := collection.Find(bson.M{"documento": numero}).One(&cliente)
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

// Consulta un cliente por correo en la base de datos y si existe retorna su conjunto de datos
func GetClientePorCorreo(ginContext *gin.Context) {
	correo := ginContext.Params.ByName("correo")
	token := ginContext.Params.ByName("token")
	fmt.Println(correo)
	if auth(token){
		session := connect();
		defer session.Close()
		collection := session.DB("my_bank_db").C("Clientes")

		cliente := Cliente{}
		err := collection.Find(bson.M{"correo": correo}).One(&cliente)
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

// Consulta el ejecutivo de un cliente
func GetEjecutivoPorCorreoCliente(ginContext *gin.Context){
	correo := ginContext.Params.ByName("correo")
	token := ginContext.Params.ByName("token")
	fmt.Println(correo)
	if auth(token){
		session := connect();
		defer session.Close()
		collection := session.DB("my_bank_db").C("Clientes")

		cliente := Cliente{}
		err := collection.Find(bson.M{"correo": correo}).One(&cliente)
		if err != nil {
			ginContext.JSON(404, gin.H{
				"auth":	"permiso concedido",
				"error":  "registro no encontrado",
				})
		}else{
			session := connect();
			defer session.Close()
			ejecutivo := Ejecutivo{}
			collection := session.DB("my_bank_db").C("Ejecutivos")
			err := collection.Find(bson.M{"id_ejecutivo": cliente.EjecutivoEncargado}).One(&ejecutivo)
			if err != nil {
				ginContext.JSON(404, gin.H{
					"auth":	"permiso concedido",
					"error":  "registro no encontrado",
					})
			}else{
				ginContext.JSON(200, ejecutivo)
			}
		}

	}else {
		ginContext.JSON(404, gin.H{
			"error":  "permiso denegado",
			})
	}
}

// Conecta a la base de datos
// Definir variable de entorno en Heroku: Settings, Config Vars
// Definir variable de entorno local: echo "export MONGO_URL=mongodb://goteam:goteam@ds019471.mlab.com:19471/my_bank_db" >> ~/.bashrc
func connect() (session *mgo.Session) {
	connectURL := os.Getenv("MONGO_URL")
	fmt.Println(connectURL)
	session, err := mgo.Dial(connectURL)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
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
