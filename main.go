package main

import (
	"database/sql"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/urwork/throw"
)

func main() {
	home, err := os.UserHomeDir()
	Throw(err)

	home = path.Join(home, ".config", "c1du")

	err = os.MkdirAll(home, os.ModePerm)
	Throw(err)

	db, err := sql.Open("sqlite3", path.Join(home, "c1du.db"))
	Throw(err)

	println(db)
}
