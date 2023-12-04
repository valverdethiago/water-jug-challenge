package restapi

import (
	"context"
	"fmt"
	"github.com/valverdethiago/water-jug-challenge/code/backend/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type WaterJugApp struct {
	serverMux *http.ServeMux
	server    *http.Server
	service   service.WaterJugService
}

func NewWaterJugServer(port int, service service.WaterJugService) WaterJugApp {
	serverMux := *http.NewServeMux()
	server := http.Server{Handler: &serverMux, Addr: fmt.Sprintf(":%d", port)}
	return WaterJugApp{
		serverMux: &serverMux,
		server:    &server,
		service:   service,
	}
}

func (app *WaterJugApp) BindEndpoints() {
	app.serverMux.HandleFunc("/", app.routingHandler)
}

func (app *WaterJugApp) Run() error {
	idleChan := make(chan struct{})
	go func(app *WaterJugApp, idleChan chan struct{}) {
		signChan := make(chan os.Signal, 1)
		signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)
		sig := <-signChan
		log.Println("shutdown:", sig)

		app.Stop(5 * time.Second)

		// Actual shutdown trigger.
		close(idleChan)
	}(app, idleChan)
	return app.server.ListenAndServe()
}

func (app *WaterJugApp) Stop(t time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()
	err := app.server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}
}
func (app *WaterJugApp) routingHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() == "/" || r.URL.RequestURI() == "" {
		if r.Method == http.MethodPost {
			solveChallengeHandler(app.service).ServeHTTP(w, r)
			return
		}
		if r.Method == http.MethodOptions {
			corsHandler().ServeHTTP(w, r)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte(`{"message": "not found"}`))
	if err != nil {
		log.Println(err.Error())
	}
}
