package steps_test

import (
	"testing"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/steps"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/utils"
)

var noHouseScenarios = []utils.TestingScenario{
	{
		About: "NoHouse -> House == nil",
		UserInfo: model.UserPersonalInformation{
			House: nil,
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{
			Home: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{
			Home: model.StepResult{
				Risk:        0,
				Eligibility: false,
			},
		},
	},
	{
		About: "NoHouse -> House != nil",
		UserInfo: model.UserPersonalInformation{
			House: &model.House{
				OwnershipStatus: "owned",
			},
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{
			Home: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{
			Home: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
}

func TestStepNoHouse(t *testing.T) {
	for _, scenario := range noHouseScenarios {
		t.Run(scenario.About, func(t *testing.T) {
			steps.NoHouse(scenario.UserInfo, scenario.InsuranceSteps)
			if scenario.InsuranceSteps.Home.Eligibility != scenario.Expected.Home.Eligibility {
				t.Errorf("Expected %v, got %v", scenario.Expected.Home.Eligibility, scenario.InsuranceSteps.Home.Eligibility)
			}
		})
	}
}

var mortgagedHouseScenarios = []utils.TestingScenario{
	{
		About: "MortgagedHouse -> House.OwnershipStatus == mortgaged",
		UserInfo: model.UserPersonalInformation{
			House: &model.House{
				OwnershipStatus: "mortgaged",
			},
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{
			Home: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{
			Home: model.StepResult{
				Risk:        1,
				Eligibility: true,
			},
		},
	},
	{
		About: "MortgagedHouse -> House.OwnershipStatus != mortgaged",
		UserInfo: model.UserPersonalInformation{
			House: &model.House{
				OwnershipStatus: "owned",
			},
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{
			Home: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{
			Home: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
}

func TestStepMortgagedHouse(t *testing.T) {
	for _, scenario := range mortgagedHouseScenarios {
		t.Run(scenario.About, func(t *testing.T) {
			steps.MortgagedHouse(scenario.UserInfo, scenario.InsuranceSteps)
			if scenario.InsuranceSteps.Home.Risk != scenario.Expected.Home.Risk {
				t.Errorf("Expected %v, got %v", scenario.Expected.Home.Risk, scenario.InsuranceSteps.Home.Risk)
			}
		})
	}
}
