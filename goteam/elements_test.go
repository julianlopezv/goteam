package goteam

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestValidateMensaje(t *testing.T) {
	
	de := "de"
	para := "para"
	mensaje := "mensaje"
	estado := "nuevo"
	fecha := "fecha"
	
	deT := Mensaje{De: de, Para: para, Fecha: fecha, Estado: estado, Mensaje: mensaje}.De
	
	assert.Equal(t, deT, de, "La estructura Mensaje esta bien definida")
}

func TestValidateProducto(t *testing.T) {
	
	nombre := "nombre"
	
	nombreT := Producto{Nombre: nombre}.Nombre
	
	assert.Equal(t, nombreT, nombre, "La estructura Producto esta bien definida")
}

func TestValidateCliente(t *testing.T) {
	
	NombreCompleto := "NombreCompleto"
	
	NombreCompletoT := Cliente{NombreCompleto: NombreCompleto}.NombreCompleto
	
	assert.Equal(t, NombreCompletoT, NombreCompleto, "La estructura Cliente esta bien definida")
}

func TestValidateEjecutivo(t *testing.T) {
	
	Nombre := "Nombre"
	
	NombreT := Ejecutivo{Nombre: Nombre}.Nombre
	
	assert.Equal(t, NombreT, Nombre, "La estructura Ejecutivo esta bien definida")
}