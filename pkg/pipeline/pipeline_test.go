package pipeline_test

import (
	"testing"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/steps"
)

func TestNewPipeline(t *testing.T) {
	pipeline := pipeline.NewPipeline()
	if pipeline == nil {
		t.Error("Expected a pipeline, got nil")
	}
	if pipeline != nil && len(pipeline.RiskSteps) != 0 {
		t.Errorf("Expected a pipeline with 0 steps, got %d", len(pipeline.RiskSteps))
	}
}

func TestAddStep(t *testing.T) {
	pipeline := pipeline.NewPipeline()
	pipeline.AddStep(steps.NoHouse)
	if pipeline != nil && len(pipeline.RiskSteps) != 1 {
		t.Errorf("Expected a pipeline with 1 step, got %d", len(pipeline.RiskSteps))
	}
}

func TestAddSteps(t *testing.T) {
	pipeline := pipeline.NewPipeline()
	pipeline.AddSteps(steps.NoHouse, steps.NoIncome)
	if pipeline != nil && len(pipeline.RiskSteps) != 2 {
		t.Errorf("Expected a pipeline with 2 steps, got %d", len(pipeline.RiskSteps))
	}
}
