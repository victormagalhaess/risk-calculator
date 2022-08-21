package riskSteps_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/riskSteps"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"
)

var noIncomeScenarios = []riskSteps.TestingScenario{
	{
		About: "NoIncome -> Income <= 0",
		UserInfo: types.UserPersonalInformation{
			Income: 0,
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Disability: types.Step{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Disability: types.Step{
				Risk:        0,
				Eligibility: false,
			},
		},
	},
	{
		About: "NoIncome -> Income > 0",
		UserInfo: types.UserPersonalInformation{
			Income: 1,
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Disability: types.Step{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Disability: types.Step{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
}

func TestNoIncome(t *testing.T) {
	for _, scenario := range noIncomeScenarios {
		t.Run(scenario.About, func(t *testing.T) {
			riskSteps.NoIncome(scenario.UserInfo, scenario.InsuranceSteps)
			if scenario.InsuranceSteps.Disability.Eligibility != scenario.Expected.Disability.Eligibility {
				t.Errorf("Expected %v, got %v", scenario.Expected.Disability.Eligibility, scenario.InsuranceSteps.Disability.Eligibility)
			}
		})
	}
}

var incomeOver200kScenarios = []riskSteps.TestingScenario{
	{
		About: "IncomeOver200k -> Income > 200000",
		UserInfo: types.UserPersonalInformation{
			Income: 200001,
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{
			Disability: types.Step{
				Risk:        0,
				Eligibility: true,
			},
			Auto: types.Step{
				Risk:        0,
				Eligibility: true,
			},
			Home: types.Step{
				Risk:        0,
				Eligibility: true,
			},
			Life: types.Step{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Disability: types.Step{
				Risk:        -1,
				Eligibility: true,
			},
			Auto: types.Step{
				Risk:        -1,
				Eligibility: true,
			},
			Home: types.Step{
				Risk:        -1,
				Eligibility: true,
			},
			Life: types.Step{
				Risk:        -1,
				Eligibility: true,
			},
		},
	},
	{
		About: "IncomeOver200k -> Income <= 200000",
		UserInfo: types.UserPersonalInformation{
			Income: 200000,
		},
		InsuranceSteps: &types.UserInsuranceAnalysisSteps{

			Disability: types.Step{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: types.UserInsuranceAnalysisSteps{
			Disability: types.Step{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
}

func TestIncomeOver200k(t *testing.T) {
	for _, scenario := range incomeOver200kScenarios {
		t.Run(scenario.About, func(t *testing.T) {
			riskSteps.IncomeOver200k(scenario.UserInfo, scenario.InsuranceSteps)
			if diff := cmp.Diff(&scenario.Expected, scenario.InsuranceSteps); diff != "" {
				t.Fatalf("Result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
