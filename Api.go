package main

import (
"./goteam"
"github.com/gin-gonic/gin"
"github.com/itsjamie/gin-cors"
"gopkg.in/mgo.v2/bson"
"strconv"
"time"
)

// Define las rutas de la API y la ejecuta
func main() {
	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, goteam.Authorization, Content-Type",
		ExposedHeaders: "*",
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
		v1.GET("/mensajes/:correo/:token", GetMensajesNuevosPorCorreoCliente)
		v1.GET("/mensaje/:de/:para/:mensaje/:token", SetNuevoMensaje)
		// v1.POST("/usuarios", PostUser)
		// v1.PUT("/usuarios/:id", UpdateUser)
		// v1.DELETE("/usuarios/:id", DeleteUser)
	}

	/*para correr en un puerto local*/
	r.Run(":1337")
	//r.Run()
}

func ImOk(ginContext *gin.Context) {
	ginContext.JSON(200, gin.H{
		"status":  "Im OK!!!",
		})
}

// Consulta la base de datos y retorna toda la coleccion de clientes
func GetMensajesNuevos(ginContext *gin.Context) {
	token := ginContext.Params.ByName("token")

	mensajes := []goteam.Mensaje{}
	
	httpStatus, obj := goteam.Select(token, nil, mensajes)
	ginContext.JSON(httpStatus, obj)
}

func GetMensajesNuevosPorCorreoCliente(ginContext *gin.Context) {
	token := ginContext.Params.ByName("token")
	correo := ginContext.Params.ByName("correo")

	mensajes := []goteam.Mensaje{}
	httpStatus, obj := goteam.Select(token, bson.M{"para": correo, "estado":"nuevo"}, mensajes)
	ginContext.JSON(httpStatus, obj)
	
}

func SetNuevoMensaje(ginContext *gin.Context) {
	token := ginContext.Params.ByName("token")
	de := ginContext.Params.ByName("de")
	para := ginContext.Params.ByName("para")
	mensaje := ginContext.Params.ByName("mensaje")
	estado := "nuevo"
	fecha := "fecha"

	httpStatus, obj := goteam.Insert(token, goteam.Mensaje{De: de, Para: para, Fecha: fecha, Estado: estado, Mensaje: mensaje})
	ginContext.JSON(httpStatus, obj)
}

// Consulta la base de datos y retorna toda la coleccion de clientes
func GetClientes(ginContext *gin.Context) {
	token := ginContext.Params.ByName("token")

	clientes := []goteam.Cliente{}
		
	httpStatus, obj := goteam.Select(token, nil, clientes)
	ginContext.JSON(httpStatus, obj)
}

// Consulta un documento de cliente en la base de datos y si existe retorna su conjunto de datos
func GetCliente(ginContext *gin.Context) {
	numeroDocumento := ginContext.Params.ByName("documento")
	numero, _ := strconv.ParseInt(numeroDocumento, 0, 64)
	token := ginContext.Params.ByName("token")

	cliente := goteam.Cliente{}
	
	httpStatus, obj := goteam.SelectOne(token, bson.M{"documento": numero}, cliente)
	ginContext.JSON(httpStatus, obj)
}

// Consulta un cliente por correo en la base de datos y si existe retorna su conjunto de datos
func GetClientePorCorreo(ginContext *gin.Context) {
	correo := ginContext.Params.ByName("correo")
	token := ginContext.Params.ByName("token")

	cliente := goteam.Cliente{}
	
	httpStatus, obj := goteam.SelectOne(token, bson.M{"correo": correo}, cliente)
	ginContext.JSON(httpStatus, obj)
}

// Consulta el ejecutivo de un cliente
func GetEjecutivoPorCorreoCliente(ginContext *gin.Context){
	correo := ginContext.Params.ByName("correo")
	token := ginContext.Params.ByName("token")
	
	cliente := goteam.Cliente{}
	
	goteam.SelectOne(token, bson.M{"correo": correo}, &cliente)
	
	ejecutivo := goteam.Ejecutivo{}
	
	httpStatus, obj := goteam.SelectOne(token, bson.M{"id_ejecutivo": cliente.EjecutivoEncargado}, ejecutivo)
	ginContext.JSON(httpStatus, obj)
}