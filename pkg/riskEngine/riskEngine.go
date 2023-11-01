package riskEngine

import (
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/riskSteps"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"
)

var Pipeline = riskSteps.NewPipeline()

func InitializePipeline() {
	Pipeline.AddSteps(
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
	)
}

func ExecutePipeline(userInfo types.UserPersonalInformation,
	insuranceSteps *types.UserInsuranceAnalysisSteps) {
	for _, step := range Pipeline.RiskSteps {
		step(userInfo, insuranceSteps)
	}

}
