package goteam

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var tokenF = "tDXNANALWtLKXxpZ8qvS3AyLM5RxwKxw1tNUebmD"

func TestAuth(t *testing.T) {
	
	assert.Equal(t, Auth(tokenF), true, "Al autenticar el token real el acceso se da")
}

func TestAuthWithoutToken(t *testing.T) {
	
	assert.Equal(t, Auth("abcd"), false, "Al autenticar el token falso el acceso se rechaza")
}