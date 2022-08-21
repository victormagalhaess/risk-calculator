package riskSteps

import (
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"
)

func AddToAllValues(insuranceSteps *types.UserInsuranceAnalysisSteps, value int) {
	insuranceSteps.Auto.Risk += value
	insuranceSteps.Disability.Risk += value
	insuranceSteps.Home.Risk += value
	insuranceSteps.Life.Risk += value
}

type TestingScenario struct {
	About          string
	UserInfo       types.UserPersonalInformation
	InsuranceSteps *types.UserInsuranceAnalysisSteps
	Expected       types.UserInsuranceAnalysisSteps
}
