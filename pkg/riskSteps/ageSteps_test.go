package riskSteps_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/riskSteps"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"
)

var over60Scenario = []riskSteps.TestingScenario{
	{
		About: "Over 60 -> Age > 60",
		UserInfo: types.UserPersonalInformation{
			Age: 61,
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Disability: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Disability: types.StepResult{
				Risk:        0,
				Eligibility: false,
			},
		},
	},
	{
		About: "Over 60 -> Age <= 60",
		UserInfo: types.UserPersonalInformation{
			Age: 60,
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Disability: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Disability: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
}

func TestOver60(t *testing.T) {
	for _, scenario := range over60Scenario {
		t.Run(scenario.About, func(t *testing.T) {
			riskSteps.Over60(scenario.UserInfo, scenario.InsuranceSteps)
			if scenario.InsuranceSteps.Disability.Eligibility != scenario.Expected.Disability.Eligibility {
				t.Errorf("Expected %v, got %v", scenario.Expected.Disability.Eligibility, scenario.InsuranceSteps.Disability.Eligibility)
			}
		})
	}
}

var under40Scenario = []riskSteps.TestingScenario{
	{
		About: "Under 40 -> Age < 40",
		UserInfo: types.UserPersonalInformation{
			Age: 39,
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Disability: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Auto: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Home: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Life: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Disability: types.StepResult{
				Risk:        -1,
				Eligibility: true,
			},
			Auto: types.StepResult{
				Risk:        -1,
				Eligibility: true,
			},
			Home: types.StepResult{
				Risk:        -1,
				Eligibility: true,
			},
			Life: types.StepResult{
				Risk:        -1,
				Eligibility: true,
			},
		},
	},
	{
		About: "Under 40 -> Age >= 40",
		UserInfo: types.UserPersonalInformation{
			Age: 40,
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Disability: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Auto: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Home: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Life: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Disability: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Auto: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Home: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Life: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
}

func TestUnder40(t *testing.T) {
	for _, scenario := range under40Scenario {
		t.Run(scenario.About, func(t *testing.T) {
			riskSteps.Under40(scenario.UserInfo, scenario.InsuranceSteps)
			if diff := cmp.Diff(&scenario.Expected, scenario.InsuranceSteps); diff != "" {
				t.Fatalf("Result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

var under30Scenario = []riskSteps.TestingScenario{
	{
		About: "Under 30 -> Age < 30",
		UserInfo: types.UserPersonalInformation{
			Age: 29,
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Disability: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Auto: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Home: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Life: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Disability: types.StepResult{
				Risk:        -1,
				Eligibility: true,
			},
			Auto: types.StepResult{
				Risk:        -1,
				Eligibility: true,
			},
			Home: types.StepResult{
				Risk:        -1,
				Eligibility: true,
			},
			Life: types.StepResult{
				Risk:        -1,
				Eligibility: true,
			},
		},
	},
	{
		About: "Under 30 -> Age >= 30",
		UserInfo: types.UserPersonalInformation{
			Age: 30,
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Disability: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Auto: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Home: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Life: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Disability: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Auto: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Home: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Life: types.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
}

func TestUnder30(t *testing.T) {
	for _, scenario := range under30Scenario {
		t.Run(scenario.About, func(t *testing.T) {
			riskSteps.Under30(scenario.UserInfo, scenario.InsuranceSteps)
			if diff := cmp.Diff(&scenario.Expected, scenario.InsuranceSteps); diff != "" {
				t.Fatalf("Result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
