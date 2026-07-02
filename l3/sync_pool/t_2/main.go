package main

import (
	"server"
)

func main() {
	myServer := server.NewServer()
	myServer.Init()
	myServer.Launch(8080)
}
