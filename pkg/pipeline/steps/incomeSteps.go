package steps

import (
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/utils"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"
)

func NoIncome(userInfo types.UserPersonalInformation,
	insuranceSteps *types.UserInsuranceAnalysisSteps) {
	if userInfo.Income <= 0 {
		insuranceSteps.Disability.Eligibility = false
	}
}

func IncomeOver200k(userInfo types.UserPersonalInformation,
	insuranceSteps *types.UserInsuranceAnalysisSteps) {
	if userInfo.Income > 200000 {
		utils.AddToAllValues(insuranceSteps, -1)
	}
}
