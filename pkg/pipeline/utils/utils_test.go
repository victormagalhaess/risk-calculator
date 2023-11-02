package utils_test

import (
	"testing"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/utils"
)

func TestAddToAllValues(t *testing.T) {
	t.Run("Positive Increment -> should all be 1", func(t *testing.T) {
		insuranceSteps := model.UserInsuranceAnalysisSteps{
			Life: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
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
		}

		utils.AddToAllValues(&insuranceSteps, 1)
		if insuranceSteps.Life.Risk != 1 {
			t.Errorf("Life.Risk is not 1")
		}
		if insuranceSteps.Disability.Risk != 1 {
			t.Errorf("Disability.Risk is not 1")
		}
		if insuranceSteps.Auto.Risk != 1 {
			t.Errorf("Auto.Risk is not 1")
		}
		if insuranceSteps.Home.Risk != 1 {
			t.Errorf("Home.Risk is not 1")
		}
	})

	t.Run("Negative Increment -> should all be -2", func(t *testing.T) {
		insuranceSteps := model.UserInsuranceAnalysisSteps{
			Life: model.StepResult{
				Risk:        0,
				Eligibility: true,
			},
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
		}

		utils.AddToAllValues(&insuranceSteps, -2)
		if insuranceSteps.Life.Risk != -2 {
			t.Errorf("Life.Risk is not -2")
		}
		if insuranceSteps.Disability.Risk != -2 {
			t.Errorf("Disability.Risk is not -2")
		}
		if insuranceSteps.Auto.Risk != -2 {
			t.Errorf("Auto.Risk is not -2")
		}
		if insuranceSteps.Home.Risk != -2 {
			t.Errorf("Home.Risk is not -2")
		}
	})

}
