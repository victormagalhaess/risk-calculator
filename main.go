package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/victormagalhaess/origin-backend-take-home-assignment/docs"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/api"
)

// @title Origin Backend Take Home Assignment
// @version 1.0
// @description This is the backend application for the Origin Backend Take Home Assignment.
// @termsOfService http://swagger.io/terms/

// @contact.name Victor Magalhães
// @contact.email hello@victordias.dev

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080

func main() {
	//parse the port flag from the command line
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	wait := time.Second * 10 //default time to wait connections to close before shutdown

	application := api.Application{}
	application.InitializeRouter()
	application.InitializeSwagger()

	// setting connection timeouts for http server
	server := &http.Server{
		Addr:         ":" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      application.Router,
	}

	// starting http server in a goroutine so that it won't block
	go func() {
		log.Println("Starting server on port " + port)
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
