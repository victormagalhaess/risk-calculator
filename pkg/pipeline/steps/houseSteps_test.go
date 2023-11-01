package steps_test

import (
	"testing"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/steps"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/utils"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"
)

var noHouseScenarios = []utils.TestingScenario{
	{
		About: "NoHouse -> House == nil",
		UserInfo: types.UserPersonalInformation{
			House: nil,
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Home: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Home: types.StepResult{
				Risk:        0,
				Eligibility: false,
			},
		},
	},
	{
		About: "NoHouse -> House != nil",
		UserInfo: types.UserPersonalInformation{
			House: &types.House{
				OwnershipStatus: "owned",
			},
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Home: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Home: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
}

func TestNoHouse(t *testing.T) {
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
		UserInfo: types.UserPersonalInformation{
			House: &types.House{
				OwnershipStatus: "mortgaged",
			},
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Home: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Home: types.StepResult{
				Risk:        1,
				Eligibility: true,
			},
		},
	},
	{
		About: "MortgagedHouse -> House.OwnershipStatus != mortgaged",
		UserInfo: types.UserPersonalInformation{
			House: &types.House{
				OwnershipStatus: "owned",
			},
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Home: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Home: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
}

func TestMortgagedHouse(t *testing.T) {
	for _, scenario := range mortgagedHouseScenarios {
		t.Run(scenario.About, func(t *testing.T) {
			steps.MortgagedHouse(scenario.UserInfo, scenario.InsuranceSteps)
			if scenario.InsuranceSteps.Home.Risk != scenario.Expected.Home.Risk {
				t.Errorf("Expected %v, got %v", scenario.Expected.Home.Risk, scenario.InsuranceSteps.Home.Risk)
			}
		})
	}
}
