package key

import (
	"github.com/u6du/ex"
	"github.com/u6du/go-rfc1924/base85"
	"golang.org/x/crypto/ed25519"
)

var GodPublic ed25519.PublicKey

func init() {
	key := "hYZUEfyAvgoJD27QXXJGpYUcU1VvAWKnQWra9LbH"
	keyByte, err := base85.DecodeString(key)
	ex.Panic(err)
	GodPublic = ed25519.PublicKey(keyByte)
}
