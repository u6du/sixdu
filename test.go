package main

import (
	"crypto/rand"
	"log"

	"github.com/u6du/ex"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ed25519"
)

func main() {
	/*
		priv := "e06d3183d14159228433ed599221b80bd0a5ce8352e4bdf0262f76786ef1c74db7e7a9fea2c0eb269d61e3b38e450a22e754941ac78479d6c54e1faf6037881d"
		pub := "77ff84905a91936367c01360803104f92432fcd904a43511876df5cdf3e7e548"
		sig := "6834284b6b24c3204eb2fea824d82f88883a3d95e8b4a21b8c0ded553d17d17ddf9a8a7104b1258f30bed3787e6cb896fca78c58f8e03b5f18f14951a87d9a08"
		// d := hex.EncodeToString([]byte(priv))
		priv := ed25519.GenerateKey(rand.Reader)

		privb, _ := hex.DecodeString(priv)
		pvk := ed25519.PrivateKey(privb)
	*/
	hasher, err := blake2b.New256(nil)

	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	ex.Panic(err)

	buffer := []byte("4:salt6:foobar3:seqi1e1:v12:Hello World!")
	sigb := ed25519.Sign(priv, buffer)
	println(len(sigb), " sign len")

	log.Println(ed25519.Verify(pub, buffer, sigb))
	log.Printf("%x\n", priv.Public())
	log.Printf("%x\n", hasher.Sum(pub))
	log.Printf("%x\n", sigb)
}
