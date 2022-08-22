package steps_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/steps"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/utils"
)

var noIncomeScenarios = []utils.TestingScenario{
	{
		About: "NoIncome -> Income <= 0",
		UserInfo: model.UserPersonalInformation{
			Income: 0,
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{
			Disability: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{
			Disability: model.StepResult{
				Risk:        0,
				Eligibility: false,
			},
		},
	},
	{
		About: "NoIncome -> Income > 0",
		UserInfo: model.UserPersonalInformation{
			Income: 1,
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{
			Disability: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{
			Disability: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
}

func TestNoIncome(t *testing.T) {
	for _, scenario := range noIncomeScenarios {
		t.Run(scenario.About, func(t *testing.T) {
			steps.NoIncome(scenario.UserInfo, scenario.InsuranceSteps)
			if scenario.InsuranceSteps.Disability.Eligibility != scenario.Expected.Disability.Eligibility {
				t.Errorf("Expected %v, got %v", scenario.Expected.Disability.Eligibility, scenario.InsuranceSteps.Disability.Eligibility)
			}
		})
	}
}

var incomeOver200kScenarios = []utils.TestingScenario{
	{
		About: "IncomeOver200k -> Income > 200000",
		UserInfo: model.UserPersonalInformation{
			Income: 200001,
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{
			Disability: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Auto: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Home: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
			Life: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{
			Disability: model.StepResult{
				Risk:        -1,
				Eligibility: true,
			},
			Auto: model.StepResult{
				Risk:        -1,
				Eligibility: true,
			},
			Home: model.StepResult{
				Risk:        -1,
				Eligibility: true,
			},
			Life: model.StepResult{
				Risk:        -1,
				Eligibility: true,
			},
		},
	},
	{
		About: "IncomeOver200k -> Income <= 200000",
		UserInfo: model.UserPersonalInformation{
			Income: 200000,
		},
		InsuranceSteps: &model.UserInsuranceAnalysisSteps{

			Disability: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
		Expected: model.UserInsuranceAnalysisSteps{
			Disability: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
		},
	},
}

func TestIncomeOver200k(t *testing.T) {
	for _, scenario := range incomeOver200kScenarios {
		t.Run(scenario.About, func(t *testing.T) {
			steps.IncomeOver200k(scenario.UserInfo, scenario.InsuranceSteps)
			if diff := cmp.Diff(&scenario.Expected, scenario.InsuranceSteps); diff != "" {
				t.Fatalf("Result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
