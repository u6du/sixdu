package config

import (
	"database/sql"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
	"github.com/u6du/ex"
)

const DriverName = "sqlite3"

func Db(name, create, insert string, args ...interface{}) *sql.DB {
	name = name + "." + DriverName
	dbpath, isNew := FilepathIsNew(name)

	db, err := sql.Open(DriverName, dbpath)
	ex.Panic(err)

	if isNew {
		db.Exec(create)
		if len(args) > 0 {
			s, err := db.Prepare(insert)
			ex.Panic(err)

			for _, i := range args {
				t := reflect.TypeOf(i)
				switch t.Kind() {
				case reflect.Interface:
					li, _ := i.([]interface{})
					_, err = s.Exec(li...)
				default:
					_, err = s.Exec(i)
				}
			}
			ex.Panic(err)
		}
	}

	return db
}
