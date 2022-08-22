package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/api/services"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/api/status"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"
)

func Risk(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		status.BadRequest(w, err)
		return
	}

	userInfo := &types.UserPersonalInformation{}
	err = json.Unmarshal(body, userInfo)
	if err != nil {
		status.BadRequest(w, err)
		return
	}

	response, err := services.Risk(*userInfo)
	if err != nil {
		status.ServerError(w, err)
		return
	}

	status.Success(w, response)
}
