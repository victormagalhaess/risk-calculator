package steps

import (
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/utils"
)

func Over60(userInfo model.UserPersonalInformation,
	insuranceSteps *model.UserInsuranceAnalysisSteps) {
	if userInfo.Age > 60 {
		insuranceSteps.Disability.Eligibility = false
		insuranceSteps.Life.Eligibility = false
	}
}

func underDesiredAge(age int, userInfo model.UserPersonalInformation,
	insuranceSteps *model.UserInsuranceAnalysisSteps) {
	if userInfo.Age < age {
		utils.AddToAllValues(insuranceSteps, -1)
	}
}

func Under40(userInfo model.UserPersonalInformation,
	insuranceSteps *model.UserInsuranceAnalysisSteps) {
	underDesiredAge(40, userInfo, insuranceSteps)
}

func Under30(userInfo model.UserPersonalInformation,
	insuranceSteps *model.UserInsuranceAnalysisSteps) {
	underDesiredAge(30, userInfo, insuranceSteps)
}
