package utils

import (
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"
)

func AddToAllValues(insuranceSteps *model.UserInsuranceAnalysisSteps, value int) {
	insuranceSteps.Auto.Risk += value
	insuranceSteps.Disability.Risk += value
	insuranceSteps.Home.Risk += value
	insuranceSteps.Life.Risk += value
}

type TestingScenario struct {
	About          string
	UserInfo       model.UserPersonalInformation
	InsuranceSteps *model.UserInsuranceAnalysisSteps
	Expected       model.UserInsuranceAnalysisSteps
}
