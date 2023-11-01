package api

import (
	"github.com/gorilla/mux"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/api/controllers"
)

type Application struct {
	Router *mux.Router
}

func (a *Application) InitializeRouter() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *Application) initializeRoutes() {
	a.Router.HandleFunc("/api/v1/healthcheck", controllers.Healthcheck).Methods("GET")
	a.Router.HandleFunc("/api/v1/risk", controllers.Risk).Methods("POST")
}
