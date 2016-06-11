package goteam

import (
"github.com/zabawaba99/firego"
"os"
)

//
// Direccion de la app en Firebase
// Definir variable de entorno en Heroku: Settings, Config Vars
// Definir variable de entorno local: echo "export APP_URL=https://radiant-inferno-2748.firebaseio.com" >> ~/.bashrc
var APP_URL = os.Getenv("APP_URL")

// Autentica un token de Google en Firebase
// Argumentos:
// token: token generado por Google
func Auth(token string) bool {
	f := firego.New(APP_URL, nil)
	f.Auth(token)
	return ProcessResponse(f)
}

// Procesa la respuesta de una peticion a Firebase
// Argumentos:
// f: instancia de Firebase con que se hizo peticion
func ProcessResponse(f *firego.Firebase) bool {
	var v map[string]interface{}
	if err := f.Value(&v); err != nil {
		return false
	}
	return true
}
