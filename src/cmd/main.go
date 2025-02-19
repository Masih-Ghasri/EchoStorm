package main

import (
	"github.com/Masih-Ghasri/EchoStorm/api"
)

func main() {
	e := api.InitServer()
	api.StartServer(e)
}
