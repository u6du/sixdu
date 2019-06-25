package main

import (
	"database/sql"
	"log"
	"os"
	"path"

	"github.com/miekg/dns"

	_ "github.com/mattn/go-sqlite3"
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
