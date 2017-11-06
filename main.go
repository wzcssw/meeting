package main

import (
	"txmeeting/api"
	"txmeeting/lib"
	"txmeeting/models"

	"github.com/ivpusic/neo"
)

func main() {
	lib.Flag()
	models.InitializeDataBase()

	lib.App = neo.App()
	lib.App.Conf.App.Addr = ":" + lib.Port
	api.Load()
	lib.App.Start()
}
