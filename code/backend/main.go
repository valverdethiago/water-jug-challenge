package main

import (
	"errors"
	"flag"
	"github.com/valverdethiago/water-jug-challenge/code/backend/restapi"
	"github.com/valverdethiago/water-jug-challenge/code/backend/service"
	"log"
	"net/http"
)

func main() {

	var port = flag.Int("port", 8080, "port")
	flag.Parse()

	waterJugService := service.NewWaterJugServiceImpl()
	server := restapi.NewWaterJugServer(*port, waterJugService)
	server.BindEndpoints()
	if err := server.Run(); !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
	log.Println("shutdown: completed")
}
