package riskSteps_test

import (
	"testing"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/riskSteps"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"
)

var dependentsScenario = []riskSteps.TestingScenario{
	{
		About: "No dependents -> 0 dependents",
		UserInfo: types.UserPersonalInformation{
			Dependents: 0,
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Life: types.Step{
				Risk:        0,
				Eligibility: true,
			},
			Disability: types.Step{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{

			Life: types.Step{
				Risk:        0,
				Eligibility: true,
			},
			Disability: types.Step{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
	{
		About: "No dependents -> dependents > 0",
		UserInfo: types.UserPersonalInformation{
			Dependents: 1,
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Life: types.Step{
				Risk:        0,
				Eligibility: true,
			},
			Disability: types.Step{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Life: types.Step{
				Risk:        1,
				Eligibility: true,
			},
			Disability: types.Step{
				Risk:        1,
				Eligibility: true,
			},
		},
	},
}

func TestDependents(t *testing.T) {
	for _, scenario := range dependentsScenario {
		t.Run(scenario.About, func(t *testing.T) {
			riskSteps.Dependents(scenario.UserInfo, scenario.InsuranceSteps)
			if scenario.InsuranceSteps.Life.Risk != scenario.Expected.Life.Risk {
				t.Errorf("Expected %v, got %v", scenario.Expected.Life.Risk, scenario.InsuranceSteps.Life.Risk)
			}
			if scenario.InsuranceSteps.Disability.Risk != scenario.Expected.Disability.Risk {
				t.Errorf("Expected %v, got %v", scenario.Expected.Disability.Risk, scenario.InsuranceSteps.Disability.Risk)
			}
		})
	}
}

var marriedScenario = []riskSteps.TestingScenario{
	{
		About: "Married -> married",
		UserInfo: types.UserPersonalInformation{
			MaritalStatus: "married",
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Life: types.Step{
				Risk:        0,
				Eligibility: true,
			},
			Disability: types.Step{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Life: types.Step{
				Risk:        1,
				Eligibility: true,
			},
			Disability: types.Step{
				Risk:        -1,
				Eligibility: true,
			},
		},
	},
	{
		About: "Married -> single",
		UserInfo: types.UserPersonalInformation{
			MaritalStatus: "single",
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Life: types.Step{
				Risk:        0,
				Eligibility: true,
			},
			Disability: types.Step{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Life: types.Step{
				Risk:        0,
				Eligibility: true,
			},
			Disability: types.Step{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
}

func TestMarried(t *testing.T) {
	for _, scenario := range marriedScenario {
		t.Run(scenario.About, func(t *testing.T) {
			riskSteps.Married(scenario.UserInfo, scenario.InsuranceSteps)
			if scenario.InsuranceSteps.Life.Risk != scenario.Expected.Life.Risk {
				t.Errorf("Expected %v, got %v", scenario.Expected.Life.Risk, scenario.InsuranceSteps.Life.Risk)
			}
			if scenario.InsuranceSteps.Disability.Risk != scenario.Expected.Disability.Risk {
				t.Errorf("Expected %v, got %v", scenario.Expected.Disability.Risk, scenario.InsuranceSteps.Disability.Risk)
			}
		})
	}
}
