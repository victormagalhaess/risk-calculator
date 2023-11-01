package services

import (
	"encoding/json"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/riskEngine"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"
)

func Risk(userInfo types.UserPersonalInformation) ([]byte, error) {
	baseRisk := userInfo.BaseRisk()
	insuranceSteps := &types.UserInsuranceAnalysisSteps{
		Disability: types.StepResult{
			Risk:        baseRisk,
			Eligibility: true,
		},
		Auto: types.StepResult{
			Risk:        baseRisk,
			Eligibility: true,
		},
		Home: types.StepResult{
			Risk:        baseRisk,
			Eligibility: true,
		},
		Life: types.StepResult{
			Risk:        baseRisk,
			Eligibility: true,
		},
	}
	riskEngine.InitializePipeline()
	riskEngine.ExecutePipeline(userInfo, insuranceSteps)
	response, err := json.Marshal(insuranceSteps.MapInsuranceAnalisysToRiskProfile())
	if err != nil {
		return nil, err
	}
	return response, nil
}
