package services_test

//This file is just a extension to the risk test suite.
//It runs the same test suite but with a few extra inputs combinations.
//It aims to ease up the development process. allowing the developer
//to test the result of the risk calculation with different inputs.
//Be aware that this test suites runs over all the test steps, so
//adding new steps will reflect in the test results.
//For the sake of simplicity, the test suites does not cover all the combinations
//but instead a few random scenarios aiming to combine different steps.

import (
	"testing"
	"time"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/api/services"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"
)

var testRiskScenarios = []struct {
	userInfo model.UserPersonalInformation
	output   string
}{
	{
		userInfo: model.UserPersonalInformation{
			Age:        35,
			Dependents: 2,
			House: &model.House{
				OwnershipStatus: "owned",
			},
			Income:        0,
			MaritalStatus: "married",
			RiskQuestions: []int8{0, 1, 0},
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 2,
			},
		},
		output: "{\"auto\":\"regular\",\"disability\":\"ineligible\",\"home\":\"economic\",\"life\":\"regular\"}",
	},
	{
		userInfo: model.UserPersonalInformation{
			Age:        30,
			Dependents: 2,
			House: &model.House{
				OwnershipStatus: "mortgaged",
			},
			Income:        100000,
			MaritalStatus: "married",
			RiskQuestions: []int8{1, 0, 0},
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 2,
			},
		},
		output: "{\"auto\":\"regular\",\"disability\":\"regular\",\"home\":\"regular\",\"life\":\"regular\"}",
	},
	{
		userInfo: model.UserPersonalInformation{
			Age:           61,
			Dependents:    0,
			House:         nil,
			Income:        0,
			MaritalStatus: "single",
			RiskQuestions: []int8{0, 0, 0},
			Vehicle:       nil,
		},
		output: "{\"auto\":\"ineligible\",\"disability\":\"ineligible\",\"home\":\"ineligible\",\"life\":\"ineligible\"}",
	},
	{
		userInfo: model.UserPersonalInformation{
			Age:        28,
			Dependents: 0,
			House: &model.House{
				OwnershipStatus: "owned",
			},
			Income:        250000,
			MaritalStatus: "single",
			RiskQuestions: []int8{0, 0, 0},
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 10,
			},
		},
		output: "{\"auto\":\"economic\",\"disability\":\"economic\",\"home\":\"economic\",\"life\":\"economic\"}",
	},
	{
		userInfo: model.UserPersonalInformation{
			Age:        35,
			Dependents: 0,
			House: &model.House{
				OwnershipStatus: "mortgaged",
			},
			Income:        199000,
			MaritalStatus: "single",
			RiskQuestions: []int8{0, 0, 1},
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 1,
			},
		},
		output: "{\"auto\":\"regular\",\"disability\":\"regular\",\"home\":\"regular\",\"life\":\"economic\"}",
	},
	{
		userInfo: model.UserPersonalInformation{
			Age:        25,
			Dependents: 5,
			House: &model.House{
				OwnershipStatus: "mortgaged",
			},
			Income:        7000,
			MaritalStatus: "single",
			RiskQuestions: []int8{0, 1, 0},
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 14,
			},
		},
		output: "{\"auto\":\"economic\",\"disability\":\"regular\",\"home\":\"economic\",\"life\":\"economic\"}",
	},
	{
		userInfo: model.UserPersonalInformation{
			Age:        68,
			Dependents: 0,
			House: &model.House{
				OwnershipStatus: "owned",
			},
			Income:        250000,
			MaritalStatus: "single",
			RiskQuestions: []int8{0, 0, 0},
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 10,
			},
		},
		output: "{\"auto\":\"economic\",\"disability\":\"ineligible\",\"home\":\"economic\",\"life\":\"ineligible\"}",
	},
	{
		userInfo: model.UserPersonalInformation{
			Age:        18,
			Dependents: 0,
			House: &model.House{
				OwnershipStatus: "owned",
			},
			Income:        250000,
			MaritalStatus: "single",
			RiskQuestions: []int8{0, 0, 0},
			Vehicle:       nil,
		},
		output: "{\"auto\":\"ineligible\",\"disability\":\"economic\",\"home\":\"economic\",\"life\":\"economic\"}",
	},
	{
		userInfo: model.UserPersonalInformation{
			Age:           18,
			Dependents:    0,
			House:         nil,
			Income:        250000,
			MaritalStatus: "single",
			RiskQuestions: []int8{0, 0, 0},
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 10,
			},
		},
		output: "{\"auto\":\"economic\",\"disability\":\"economic\",\"home\":\"ineligible\",\"life\":\"economic\"}",
	},
	{
		userInfo: model.UserPersonalInformation{
			Age:        29,
			Dependents: 4,
			House: &model.House{
				OwnershipStatus: "owned",
			},
			Income:        20000,
			MaritalStatus: "married",
			RiskQuestions: []int8{1, 1, 1},
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 8,
			},
		},
		output: "{\"auto\":\"regular\",\"disability\":\"regular\",\"home\":\"regular\",\"life\":\"responsible\"}",
	},
	{
		userInfo: model.UserPersonalInformation{
			Age:        55,
			Dependents: 4,
			House: &model.House{
				OwnershipStatus: "owned",
			},
			Income:        20000,
			MaritalStatus: "married",
			RiskQuestions: []int8{1, 1, 1},
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 8,
			},
		},
		output: "{\"auto\":\"responsible\",\"disability\":\"responsible\",\"home\":\"responsible\",\"life\":\"responsible\"}",
	},
	{
		userInfo: model.UserPersonalInformation{
			Age:        22,
			Dependents: 17,
			House: &model.House{
				OwnershipStatus: "owned",
			},
			Income:        20180000,
			MaritalStatus: "married",
			RiskQuestions: []int8{0, 1, 0},
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 2022,
			},
		},
		output: "{\"auto\":\"economic\",\"disability\":\"economic\",\"home\":\"economic\",\"life\":\"economic\"}",
	},
	{
		userInfo: model.UserPersonalInformation{}, //abuse of golang's variables 'zero' (default) value
		output:   "{\"auto\":\"ineligible\",\"disability\":\"ineligible\",\"home\":\"ineligible\",\"life\":\"ineligible\"}",
	},
	{
		userInfo: model.UserPersonalInformation{
			Age:        22,
			Dependents: 17,
			House: &model.House{
				OwnershipStatus: "owned",
			},
			Income:        7000,
			MaritalStatus: "married",
			RiskQuestions: []int8{0, 0, 0},
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 2022,
			},
		},
		output: "{\"auto\":\"ineligible\",\"disability\":\"ineligible\",\"home\":\"ineligible\",\"life\":\"ineligible\"}",
	},
	{
		userInfo: model.UserPersonalInformation{
			Age:        22,
			Dependents: 17,
			House: &model.House{
				OwnershipStatus: "rented",
			},
			Income:        7000,
			MaritalStatus: "married",
			RiskQuestions: []int8{0, 0, 0},
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 2022,
			},
		},
		output: "{\"auto\":\"ineligible\",\"disability\":\"ineligible\",\"renters\":\"ineligible\",\"life\":\"ineligible\"}",
	},
	{
		userInfo: model.UserPersonalInformation{
			Age:        68,
			Dependents: 0,
			House: &model.House{
				OwnershipStatus: "rented",
			},
			Income:        250000,
			MaritalStatus: "single",
			RiskQuestions: []int8{0, 0, 0},
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 10,
			},
		},
		output: "{\"auto\":\"economic\",\"disability\":\"ineligible\",\"renters\":\"economic\",\"life\":\"ineligible\"}",
	},
	{
		userInfo: model.UserPersonalInformation{
			Age:        55,
			Dependents: 4,
			House: &model.House{
				OwnershipStatus: "rented",
			},
			Income:        20000,
			MaritalStatus: "married",
			RiskQuestions: []int8{1, 1, 1},
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 8,
			},
		},
		output: "{\"auto\":\"responsible\",\"disability\":\"responsible\",\"renters\":\"responsible\",\"life\":\"responsible\"}",
	},
}

func TestRisk_When_RiskOnRegularAllButIneligibleForDisability_Then_ReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[0]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_RiskOnRegularAll_Then_ReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[1]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_RiskOnIneligibleAll_Then_ReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[2]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_RiskOnEconomicAllWithExtremeValues_Then_ReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[3]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_RiskOnDisabilityRegularRestEconomic_Then_ReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[4]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_RiskOnIneligibleLifeAndDisability_Then_ReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[5]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_RiskOnIneligibleAuto_Then_ReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[6]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_RiskOnIneligibleLifeAndDisabilityRestEconomic_Then_ReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[7]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_RiskOnHomeIneligibleRestEconomic_Then_ReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[8]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_RiskOnResponsibleLife_Then_ReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[9]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_RiskOnResponsibleAll_Then_ReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[10]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_RiskOnEconomicAll_Then_ReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[11]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_ZeroedInputRiskOnIneligibleAll_ThenReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[12]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_RiskOnIneligibleAllHomeOwned_ThenReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[13]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_RiskOnIneligibleAllHomeRented_ThenReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[14]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_RiskOnEconomicWithRentedHome_ThenReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[15]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}

func TestRisk_When_RiskOnResponsibleWithRentedHome_ThenReturnRisk(t *testing.T) {
	scenario := testRiskScenarios[16]
	response, err := services.Risk(scenario.userInfo)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if string(response) != scenario.output {
		t.Errorf("Expected data %s, got %s", scenario.output, response)
	}
}
