package initialize

import (
	"service_one/database"
	"service_one/globals"
)

func Initialize() {
	globals.Init()
	database.Init()
}
