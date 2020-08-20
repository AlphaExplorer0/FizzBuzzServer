// FizzBuzz server.
//
//     BasePath: /
//     Version: 1.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"leboncointest.com/fizzbuzzserver/internal/api"
	"leboncointest.com/fizzbuzzserver/internal/config"
)

func setup(conf *config.Config) *api.API {
	api := api.API{}
	return &api
}

/***/
func main() {

	fmt.Println("Starting FizzBuzz server")

	var err error

	// Initialize config
	config := config.Config{}
	if err = config.Init(); err != nil {
		fmt.Printf("Config initialization error : %s\n", err.Error())
	}

	// Initalize API
	api := setup(&config)

	router := mux.NewRouter()
	api.Register(router, &config)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.BindIP, config.BindPort),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("HTTP Server stopped : %s", err.Error())
		}
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
	fmt.Println("Stopping FizzBuzz server")
}
