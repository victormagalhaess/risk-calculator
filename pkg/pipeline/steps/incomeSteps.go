package steps

import (
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/utils"
)

func NoIncome(userInfo model.UserPersonalInformation,
	insuranceSteps *model.UserInsuranceAnalysisSteps) {
	if userInfo.Income <= 0 {
		insuranceSteps.Disability.Eligibility = false
	}
}

func IncomeOver200k(userInfo model.UserPersonalInformation,
	insuranceSteps *model.UserInsuranceAnalysisSteps) {
	if userInfo.Income > 200000 {
		utils.AddToAllValues(insuranceSteps, -1)
	}
}

func IncomeUnder25kNoRiskQuestions(userInfo model.UserPersonalInformation,
	insuranceSteps *model.UserInsuranceAnalysisSteps) {
	if userInfo.Income < 25000 && !userInfo.HasRiskQuestions() {
		insuranceSteps.SetEligibility(false)
	}
}
