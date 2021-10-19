package main

import (
	"transaction/initialize"
	"transaction/router"
)

func main() {
	initialize.Initialize()
	router.StartApplication()
}
