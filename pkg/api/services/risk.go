package services

import (
	"encoding/json"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/engine"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"
)

func Risk(userInfo model.UserPersonalInformation) ([]byte, error) {
	baseRisk := userInfo.BaseRisk()
	insuranceSteps := &model.UserInsuranceAnalysisSteps{
		Disability: model.StepResult{
			Risk:        baseRisk,
			Eligibility: true,
		},
		Auto: model.StepResult{
			Risk:        baseRisk,
			Eligibility: true,
		},
		Home: model.StepResult{
			Risk:        baseRisk,
			Eligibility: true,
		},
		Life: model.StepResult{
			Risk:        baseRisk,
			Eligibility: true,
		},
	}
	engine.InitializePipeline()
	engine.ExecutePipeline(userInfo, insuranceSteps)
	response, err := json.Marshal(insuranceSteps.MapInsuranceAnalisysToRiskProfile())
	if err != nil {
		return nil, err
	}
	return response, nil
}
