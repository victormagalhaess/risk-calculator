package steps_test

import (
	"testing"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/steps"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/utils"
)

var dependentsScenario = []utils.TestingScenario{
	{
		About: "No dependents -> 0 dependents",
		UserInfo: model.UserPersonalInformation{
			Dependents: 0,
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{
			Life: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Disability: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{

			Life: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Disability: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
	{
		About: "No dependents -> dependents > 0",
		UserInfo: model.UserPersonalInformation{
			Dependents: 1,
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{
			Life: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Disability: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{
			Life: model.StepResult{
				Risk:        1,
				Eligibility: true,
			},
			Disability: model.StepResult{
				Risk:        1,
				Eligibility: true,
			},
		},
	},
}

func TestDependents(t *testing.T) {
	for _, scenario := range dependentsScenario {
		t.Run(scenario.About, func(t *testing.T) {
			steps.Dependents(scenario.UserInfo, scenario.InsuranceSteps)
			if scenario.InsuranceSteps.Life.Risk != scenario.Expected.Life.Risk {
				t.Errorf("Expected %v, got %v", scenario.Expected.Life.Risk, scenario.InsuranceSteps.Life.Risk)
			}
			if scenario.InsuranceSteps.Disability.Risk != scenario.Expected.Disability.Risk {
				t.Errorf("Expected %v, got %v", scenario.Expected.Disability.Risk, scenario.InsuranceSteps.Disability.Risk)
			}
		})
	}
}

var marriedScenario = []utils.TestingScenario{
	{
		About: "Married -> married",
		UserInfo: model.UserPersonalInformation{
			MaritalStatus: "married",
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{
			Life: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Disability: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{
			Life: model.StepResult{
				Risk:        1,
				Eligibility: true,
			},
			Disability: model.StepResult{
				Risk:        -1,
				Eligibility: true,
			},
		},
	},
	{
		About: "Married -> single",
		UserInfo: model.UserPersonalInformation{
			MaritalStatus: "single",
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{
			Life: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Disability: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{
			Life: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Disability: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
}

func TestMarried(t *testing.T) {
	for _, scenario := range marriedScenario {
		t.Run(scenario.About, func(t *testing.T) {
			steps.Married(scenario.UserInfo, scenario.InsuranceSteps)
			if scenario.InsuranceSteps.Life.Risk != scenario.Expected.Life.Risk {
				t.Errorf("Expected %v, got %v", scenario.Expected.Life.Risk, scenario.InsuranceSteps.Life.Risk)
			}
			if scenario.InsuranceSteps.Disability.Risk != scenario.Expected.Disability.Risk {
				t.Errorf("Expected %v, got %v", scenario.Expected.Disability.Risk, scenario.InsuranceSteps.Disability.Risk)
			}
		})
	}
}
