package riskEngine_test

import (
	"testing"

	riskengine "github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/riskEngine"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/riskSteps"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"
)

func TestInitializePipeline(t *testing.T) {
	riskengine.InitializePipeline()
	if len(riskengine.Pipeline.RiskSteps) != 11 {
		t.Errorf("Error pipeline should have been initialized with 11 steps, but it has %d steps", len(riskengine.Pipeline.RiskSteps))
	}
	riskengine.Pipeline.RiskSteps = []riskSteps.Step{}
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
