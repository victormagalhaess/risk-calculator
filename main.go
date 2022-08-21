package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/api"
)

func main() {
	wait := time.Second * 10 //default time to wait connections to close before shutdown

	application := api.Application{}
	application.InitializeRouter()

	// setting connection timeouts for http server
	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      application.Router,
	}

	// starting http server in a goroutine so that it won't block
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// waiting for interrupt signal to gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// gracefully shutting down the server
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("Shutting down")
	os.Exit(0)
}
