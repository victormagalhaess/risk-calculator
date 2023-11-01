package services_test

import (
	"testing"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/api/services"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"
)

func TestRisk(t *testing.T) {
	userInfo := model.UserPersonalInformation{
		Age:        35,
		Dependents: 2,
		House: &model.House{
			OwnershipStatus: "owned",
		},
		Income:        0,
		MaritalStatus: "married",
		RiskQuestions: []int8{0, 1, 0},
		Vehicle: &model.Vehicle{
			Year: 2018,
		},
	}

	output := "{\"auto\":\"economic\",\"disability\":\"ineligible\",\"home\":\"economic\",\"life\":\"regular\"}"
	response, err := services.Risk(userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != output {
		t.Errorf("Expected data %s, got %s", output, response)
	}
}
