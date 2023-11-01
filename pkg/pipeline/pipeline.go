package pipeline

import "github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"

type Step func(userInfo types.UserPersonalInformation, insuranceSteps *types.UserInsuranceAnalysisSteps)
type RiskPipeline struct {
	RiskSteps []Step
}

func (pipeline *RiskPipeline) AddStep(step Step) {
	pipeline.RiskSteps = append(pipeline.RiskSteps, step)
}

func (pipeline *RiskPipeline) AddSteps(steps ...Step) {
	pipeline.RiskSteps = append(pipeline.RiskSteps, steps...)
}

// In order to add a new risk step, add a new step function in the pipeline/steps package and add it to the pipeline
// in the riskEngine package.
func NewPipeline() *RiskPipeline {
	return &RiskPipeline{
		RiskSteps: []Step{},
	}

}
