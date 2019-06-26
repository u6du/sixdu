package main

import (
	cryptoRand "crypto/rand"
	"database/sql"
	"encoding/base64"
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
	"github.com/miekg/dns"
	"github.com/prysmaticlabs/prysm/shared/bls"
	. "github.com/urwork/throw"
)


func main() {
	home, err := os.UserHomeDir()
	Throw(err)

	home = path.Join(home, ".config", "6du")

	err = os.MkdirAll(home, os.ModePerm)
	Throw(err)

	db, err := sql.Open("sqlite3", path.Join(home, "6du.db"))
	Throw(err)
	defer db.Close()

	println(db)
	token := make([]byte, 32)

	n, err := cryptoRand.Read(token)
	println("cryptoRand.Read",n)
	Throw(err)

	secret,err := bls.SecretKeyFromBytes(token)
	Throw(err)
	println("secret", base64.RawURLEncoding.Encode(secret))

	target := "6du-boot.6du.world"


	server := "8.8.8.8"

	c := dns.Client{}
	m := dns.Msg{}
	m.SetQuestion(target+".", dns.TypeTXT)
	r, t, err := c.Exchange(&m, server+":53")
	Throw(err)

	log.Printf("Took %v", t)
	if len(r.Answer) == 0 {
		log.Fatal("No results")
	}
	for _, ans := range r.Answer {
		record := ans.(*dns.TXT)
		log.Printf("%s", record.Txt)
	}

}
