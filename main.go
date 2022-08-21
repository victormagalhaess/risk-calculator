package main

import (
	"net/http"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/api"
)

func main() {
	application := api.Application{}
	application.InitializeRouter()
	http.ListenAndServe(":8080", application.Router)
}
