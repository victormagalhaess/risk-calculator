package riskSteps

import "github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"

func Dependents(userInfo types.UserPersonalInformation,
	insuranceSteps *types.UserInsuranceAnalysisSteps) {
	if userInfo.Dependents > 0 {
		insuranceSteps.Life.Risk++
		insuranceSteps.Disability.Risk++
	}
}

func Married(userInfo types.UserPersonalInformation,
	insuranceSteps *types.UserInsuranceAnalysisSteps) {
	if userInfo.IsUserMarried() {
		insuranceSteps.Life.Risk++
		insuranceSteps.Disability.Risk--
	}
}
