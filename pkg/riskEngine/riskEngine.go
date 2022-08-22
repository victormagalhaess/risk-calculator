package riskEngine

import (
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/steps"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"
)

var Pipeline = pipeline.NewPipeline()

func InitializePipeline() {
	Pipeline.AddSteps(
		steps.Over60,
		steps.Under40,
		steps.Under30,
		steps.NoHouse,
		steps.MortgagedHouse,
		steps.NoIncome,
		steps.IncomeOver200k,
		steps.Married,
		steps.Dependents,
		steps.NoVehicle,
		steps.VehicleBuiltLast5Years,
	)
}

func ExecutePipeline(userInfo types.UserPersonalInformation,
	insuranceSteps *types.UserInsuranceAnalysisSteps) {
	for _, step := range Pipeline.RiskSteps {
		step(userInfo, insuranceSteps)
	}

}
