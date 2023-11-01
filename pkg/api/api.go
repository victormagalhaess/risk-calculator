package api

import "github.com/gorilla/mux"

type Application struct {
	Router *mux.Router
}

func (a *Application) InitializeRouter() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *Application) initializeRoutes() {
	a.Router.HandleFunc("/api/v1/healthcheck", healthcheck).Methods("GET")
}
