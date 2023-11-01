package riskSteps

import "github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"

func NoHouse(userInfo types.UserPersonalInformation,
	insuranceSteps *types.UserInsuranceAnalysisSteps) {
	if userInfo.House == nil {
		insuranceSteps.Home.Eligibility = false
	}
}

func MortgagedHouse(userInfo types.UserPersonalInformation,
	insuranceSteps *types.UserInsuranceAnalysisSteps) {
	if userInfo.IsHouseMortgaged() {
		insuranceSteps.Home.Risk++
		insuranceSteps.Disability.Risk++
	}
}
