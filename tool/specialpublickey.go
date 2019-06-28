package main

import (
	"crypto/rand"
	"io/ioutil"
	"path"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/u6du/ex"
	"github.com/u6du/go-rfc1924/base85"
	"golang.org/x/crypto/ed25519"

	"sixdu/config"
)

func main() {
	count := 0
NEXT:
	for {
		_, private, err := ed25519.GenerateKey(rand.Reader)
		ex.Panic(err)
		count += 1
		if count%100000 == 0 {
			log.Info().Int("count", count).Msg("")
		}
		key := base85.EncodeToString(private.Public().(ed25519.PublicKey))

		for _, c := range "<>&`$%=-|@{}()*#;_!^?~+" {
			if strings.Index(key, string(c)) >= 0 {
				continue NEXT
			}
		}

		log.Print(key)
		filepath := config.Filepath(path.Join(config.Toml.User.Root, "key", "6du.private"))
		ioutil.WriteFile(filepath, private.Seed(), 0600)
		break
	}
}
