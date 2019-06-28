package config

import (
	"io/ioutil"

	"github.com/u6du/ex"
)

func FileByte(filename string, init func() []byte) []byte {
	filepath, isNew := FilepathIsNew(filename)
	var txt []byte
	if isNew {
		txt = init()
		ioutil.WriteFile(filepath, txt, 0600)
	} else {
		var err error
		txt, err = ioutil.ReadFile(filepath)
		ex.Panic(err)
	}
	return txt
}
