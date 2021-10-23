package initialize

import (
	"service_three/database"
	"service_three/globals"
)

func Initialize() {
	globals.Init()
	database.Init()
}
