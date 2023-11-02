package engine_test

import (
	"testing"

	engine "github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/engine"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline"
)

func TestEngine_WhenInitializePipeline_Then_ShouldBeInitializedWithSteps(t *testing.T) {
	engine.InitializePipeline()
	if len(engine.Pipeline.RiskSteps) != 12 {
		t.Errorf("Error pipeline should have been initialized with 12 steps, but it has %d steps", len(engine.Pipeline.RiskSteps))
	}
	engine.Pipeline.RiskSteps = []pipeline.Step{}
}

func TestEngine_When_ExecutePipeline_ThenShouldExecuteAllSteps(t *testing.T) {
	called := 0
	mockSteps := []pipeline.Step{}

	for i := 0; i < 10; i++ {
		mockSteps = append(mockSteps,
			func(userInfo model.UserPersonalInformation,
				insuranceSteps *model.UserInsuranceAnalysisSteps) {
				called++
			})
	}
	engine.Pipeline.AddSteps(mockSteps...)
	engine.ExecutePipeline(model.UserPersonalInformation{}, &model.UserInsuranceAnalysisSteps{})
	if called != 10 {
		t.Errorf("Error pipeline should have been called 10 times, but it was called %d times", called)
	}
}
