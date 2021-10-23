package initialize

import (
	"service_two/database"
	"service_two/globals"
)

func Initialize() {
	globals.Init()
	database.Init()
}
