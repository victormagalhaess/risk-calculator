package steps

import "github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"

func NoHouse(userInfo model.UserPersonalInformation,
	insuranceSteps *model.UserInsuranceAnalysisSteps) {
	if userInfo.House == nil {
		insuranceSteps.Home.Eligibility = false
	}
}

func MortgagedHouse(userInfo model.UserPersonalInformation,
	insuranceSteps *model.UserInsuranceAnalysisSteps) {
	if userInfo.IsHouseMortgaged() {
		insuranceSteps.Home.Risk++
		insuranceSteps.Disability.Risk++
	}
}
