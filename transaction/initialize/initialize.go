package initialize

import (
	"transaction/globals"
	"transaction/validator"
)

func Initialize() {
	globals.Init()
	validator.Init()
}
