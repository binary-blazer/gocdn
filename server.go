package main

import (
	"gocdn/config"
	"gocdn/server"
)

func main() {
	conf := config.GetConfig()

	server.StartServer(conf.Webserver.Port)
}
