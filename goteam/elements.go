package goteam

import (
"gopkg.in/mgo.v2/bson"
)

type Mensaje struct{
	Id              bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	De string `bson:"de" json:"de"`
	Para   string `bson:"para" json:"para"`
	Fecha  string `bson:"fecha" json:"fecha"`
	Estado string `bson:"estado" json:"estado"`
	Mensaje string `bson:"mensaje" json:"mensaje"`
}

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

// Estructura Ubicacion
type Ubicacion struct {
	Longitud float64 `bson:"longitud" json:"longitud"`
	Latitud float64 `bson:"latitud" json:"latitud"`
}
