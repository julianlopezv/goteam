package goteam

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

var token = "tDXNANALWtLKXxpZ8qvS3AyLM5RxwKxw1tNUebmD"

func TestSelect(t *testing.T) {
	
	_, res := Select(token, nil, nil)
	assert.Equal(t, res, bson.M{"error":  "permiso denegado"}, "Al autenticar el token abcd el acceso se rechaza")
}

func TestSelectWithoutToken(t *testing.T) {
	
	_, res := Select("abcd", nil, nil)
	assert.Equal(t, res, bson.M{"error":  "permiso denegado"}, "Al autenticar el token abcd el acceso se rechaza")
}

func TestSelectOne(t *testing.T) {
	
	_, res := SelectOne(token, nil, nil)
	assert.Equal(t, res, bson.M{"error":  "permiso denegado"}, "Al autenticar el token abcd el acceso se rechaza")
}

func TestSelectOneWithoutToken(t *testing.T) {
	
	_, res := SelectOne("abcd", nil, nil)
	assert.Equal(t, res, bson.M{"error":  "permiso denegado"}, "Al autenticar el token abcd el acceso se rechaza")
}

func TestInsertWithoutToken(t *testing.T) {
	
	_, res := Insert(token, nil)
	assert.Equal(t, res, bson.M{"error":  "permiso denegado"}, "Al autenticar el token abcd el acceso se rechaza")
}

func TestInsert(t *testing.T) {
	
	_, res := Insert("abcd", nil)
	assert.Equal(t, res, bson.M{"error":  "permiso denegado"}, "Al autenticar el token abcd el acceso se rechaza")
}