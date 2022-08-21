package riskengine_test

import (
	"testing"

	riskengine "github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/riskEngine"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/riskSteps"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"
)

func TestInitializePipeline(t *testing.T) {
	t.Run("TestInitializePipeline", func(t *testing.T) {
		pipeline := riskSteps.NewPipeline()
		if len(pipeline.RiskSteps) != 0 {
			t.Errorf("Error pipeline should be empty, but it was not")
		}

		steps := []riskSteps.Step{
			riskSteps.Over60,
			riskSteps.Under40,
			riskSteps.Under30,
			riskSteps.NoHouse,
			riskSteps.MortgagedHouse,
			riskSteps.NoIncome,
			riskSteps.IncomeOver200k,
			riskSteps.Married,
			riskSteps.Dependents,
			riskSteps.NoVehicle,
			riskSteps.VehicleBuiltLast5Years,
		}

		pipeline.AddSteps(steps...)
		if len(pipeline.RiskSteps) != len(steps) {
			t.Errorf("Error pipeline should have %d steps, but it has %d", len(steps), len(pipeline.RiskSteps))
		}

	})
}

func TestExecutePipeline(t *testing.T) {
	called := 0
	mockSteps := []riskSteps.Step{}

	for i := 0; i < 10; i++ {
		mockSteps = append(mockSteps,
			func(userInfo types.UserPersonalInformation,
				insuranceSteps *types.UserInsuranceAnalysisSteps) {
				called++
			})
	}
	riskengine.Pipeline.AddSteps(mockSteps...)
	riskengine.ExecutePipeline(types.UserPersonalInformation{}, &types.UserInsuranceAnalysisSteps{})
	if called != 10 {
		t.Errorf("Error pipeline should have been called 10 times, but it was called %d times", called)
	}
}
