package steps_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/steps"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/utils"
)

var over60Scenario = []utils.TestingScenario{
	{
		About: "Over 60 -> Age > 60",
		UserInfo: model.UserPersonalInformation{
			Age: 61,
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
		About: "Over 60 -> Age <= 60",
		UserInfo: model.UserPersonalInformation{
			Age: 60,
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

func TestStepOver60(t *testing.T) {
	for _, scenario := range over60Scenario {
		t.Run(scenario.About, func(t *testing.T) {
			steps.Over60(scenario.UserInfo, scenario.InsuranceSteps)
			if scenario.InsuranceSteps.Disability.Eligibility != scenario.Expected.Disability.Eligibility {
				t.Errorf("Expected %v, got %v", scenario.Expected.Disability.Eligibility, scenario.InsuranceSteps.Disability.Eligibility)
			}
		})
	}
}

var under40Scenario = []utils.TestingScenario{
	{
		About: "Under 40 -> Age < 40",
		UserInfo: model.UserPersonalInformation{
			Age: 39,
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
		About: "Under 40 -> Age >= 40",
		UserInfo: model.UserPersonalInformation{
			Age: 40,
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
	},
}

func TestStepUnder40(t *testing.T) {
	for _, scenario := range under40Scenario {
		t.Run(scenario.About, func(t *testing.T) {
			steps.Under40(scenario.UserInfo, scenario.InsuranceSteps)
			if diff := cmp.Diff(&scenario.Expected, scenario.InsuranceSteps); diff != "" {
				t.Fatalf("Result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

var under30Scenario = []utils.TestingScenario{
	{
		About: "Under 30 -> Age < 30",
		UserInfo: model.UserPersonalInformation{
			Age: 29,
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
		About: "Under 30 -> Age >= 30",
		UserInfo: model.UserPersonalInformation{
			Age: 30,
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
	},
}

func TestStepUnder30(t *testing.T) {
	for _, scenario := range under30Scenario {
		t.Run(scenario.About, func(t *testing.T) {
			steps.Under30(scenario.UserInfo, scenario.InsuranceSteps)
			if diff := cmp.Diff(&scenario.Expected, scenario.InsuranceSteps); diff != "" {
				t.Fatalf("Result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
