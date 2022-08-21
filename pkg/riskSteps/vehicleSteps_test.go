package riskSteps_test

import (
	"testing"
	"time"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/riskSteps"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"
)

var noVehicleScenarios = []riskSteps.TestingScenario{
	{
		About: "NoVehicle -> Vehicle == nil",
		UserInfo: types.UserPersonalInformation{
			Vehicle: nil,
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Auto: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Auto: types.StepResult{
				Risk:        0,
				Eligibility: false,
			},
		},
	},
	{
		About: "NoVehicle -> Vehicle != nil",
		UserInfo: types.UserPersonalInformation{
			Vehicle: &types.Vehicle{
				Year: 2019,
			},
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Auto: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Auto: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
}

func TestNoVehicle(t *testing.T) {
	for _, scenario := range noVehicleScenarios {
		t.Run(scenario.About, func(t *testing.T) {
			riskSteps.NoVehicle(scenario.UserInfo, scenario.InsuranceSteps)
			if scenario.InsuranceSteps.Auto.Eligibility != scenario.Expected.Auto.Eligibility {
				t.Errorf("Expected %v, got %v", scenario.Expected.Auto.Eligibility, scenario.InsuranceSteps.Auto.Eligibility)
			}
		})
	}
}

var vehicleBuiltLast5YearsScenarios = []riskSteps.TestingScenario{
	{
		About: "VehicleBuiltLast5Years -> VehicleBuiltLast5Years == true",
		UserInfo: types.UserPersonalInformation{
			Vehicle: &types.Vehicle{
				Year: time.Now().Year() - 5,
			},
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Auto: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Auto: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
	{
		About: "VehicleBuiltLast5Years -> VehicleBuiltLast5Years == false",
		UserInfo: types.UserPersonalInformation{
			Vehicle: &types.Vehicle{
				Year: time.Now().Year() - 4,
			},
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Auto: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Auto: types.StepResult{
				Risk:        1,
				Eligibility: true,
			},
		},
	},
}

func TestVehicleBuiltLast5Years(t *testing.T) {
	for _, scenario := range vehicleBuiltLast5YearsScenarios {
		t.Run(scenario.About, func(t *testing.T) {
			riskSteps.VehicleBuiltLast5Years(scenario.UserInfo, scenario.InsuranceSteps)
			if scenario.InsuranceSteps.Auto.Risk != scenario.Expected.Auto.Risk {
				t.Errorf("Expected %v, got %v", scenario.Expected.Auto.Risk, scenario.InsuranceSteps.Auto.Risk)
			}
		})
	}
}
