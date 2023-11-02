package steps_test

import (
	"testing"
	"time"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/steps"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/utils"
)

var noVehicleScenarios = []utils.TestingScenario{
	{
		About: "NoVehicle -> Vehicle == nil",
		UserInfo: model.UserPersonalInformation{
			Vehicle: nil,
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{
			Auto: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{
			Auto: model.StepResult{
				Risk:        0,
				Eligibility: false,
			},
		},
	},
	{
		About: "NoVehicle -> Vehicle != nil",
		UserInfo: model.UserPersonalInformation{
			Vehicle: &model.Vehicle{
				Year: 2019,
			},
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{
			Auto: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{
			Auto: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
}

func TestStepNoVehicle(t *testing.T) {
	for _, scenario := range noVehicleScenarios {
		t.Run(scenario.About, func(t *testing.T) {
			steps.NoVehicle(scenario.UserInfo, scenario.InsuranceSteps)
			if scenario.InsuranceSteps.Auto.Eligibility != scenario.Expected.Auto.Eligibility {
				t.Errorf("Expected %v, got %v", scenario.Expected.Auto.Eligibility, scenario.InsuranceSteps.Auto.Eligibility)
			}
		})
	}
}

var vehicleBuiltLast5YearsScenarios = []utils.TestingScenario{
	{
		About: "VehicleBuiltLast5Years -> VehicleBuiltLast5Years == true",
		UserInfo: model.UserPersonalInformation{
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 5,
			},
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{
			Auto: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{
			Auto: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
	{
		About: "VehicleBuiltLast5Years -> VehicleBuiltLast5Years == false",
		UserInfo: model.UserPersonalInformation{
			Vehicle: &model.Vehicle{
				Year: time.Now().Year() - 4,
			},
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{
			Auto: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{
			Auto: model.StepResult{
				Risk:        1,
				Eligibility: true,
			},
		},
	},
}

func TestStepVehicleBuiltLast5Years(t *testing.T) {
	for _, scenario := range vehicleBuiltLast5YearsScenarios {
		t.Run(scenario.About, func(t *testing.T) {
			steps.VehicleBuiltLast5Years(scenario.UserInfo, scenario.InsuranceSteps)
			if scenario.InsuranceSteps.Auto.Risk != scenario.Expected.Auto.Risk {
				t.Errorf("Expected %v, got %v", scenario.Expected.Auto.Risk, scenario.InsuranceSteps.Auto.Risk)
			}
		})
	}
}
