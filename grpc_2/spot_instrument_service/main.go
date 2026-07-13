package main

import (
	"fmt"
	"spot_instrument_service/server"
)

func main() {
	server := server.NewSpotInstrumentServer()

	fmt.Println(server)
}
