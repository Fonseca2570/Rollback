package globals

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

var Conf Config

func Init() {
	c, err := ioutil.ReadFile("conf.toml")
	if err != nil {
		panic(err)
	}


	if _, err = toml.Decode(string(c), &Conf); err != nil {
		panic(err)
	}

}
