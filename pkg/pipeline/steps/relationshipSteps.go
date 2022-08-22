package steps

import "github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"

func Dependents(userInfo model.UserPersonalInformation,
	insuranceSteps *model.UserInsuranceAnalysisSteps) {
	if userInfo.Dependents > 0 {
		insuranceSteps.Life.Risk++
		insuranceSteps.Disability.Risk++
	}
}

func Married(userInfo model.UserPersonalInformation,
	insuranceSteps *model.UserInsuranceAnalysisSteps) {
	if userInfo.IsUserMarried() {
		insuranceSteps.Life.Risk++
		insuranceSteps.Disability.Risk--
	}
}
