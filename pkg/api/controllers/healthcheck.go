package controllers

import (
	"net/http"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/api/status"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	status.Success(w, []byte("OK"))
}
