package main

import (
	"gingonic/src/config"
	"gingonic/src/helper"
	"gingonic/src/routes"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helper.Migration()
	defer config.DB.Close()

	routes.Router()
}