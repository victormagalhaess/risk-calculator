package steps

import (
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/pipeline/utils"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"
)

func Over60(userInfo types.UserPersonalInformation,
	insuranceSteps *types.UserInsuranceAnalysisSteps) {
	if userInfo.Age > 60 {
		insuranceSteps.Disability.Eligibility = false
		insuranceSteps.Life.Eligibility = false
	}
}

func underDesiredAge(age int, userInfo types.UserPersonalInformation,
	insuranceSteps *types.UserInsuranceAnalysisSteps) {
	if userInfo.Age < age {
		utils.AddToAllValues(insuranceSteps, -1)
	}
}

func Under40(userInfo types.UserPersonalInformation,
	insuranceSteps *types.UserInsuranceAnalysisSteps) {
	underDesiredAge(40, userInfo, insuranceSteps)
}

func Under30(userInfo types.UserPersonalInformation,
	insuranceSteps *types.UserInsuranceAnalysisSteps) {
	underDesiredAge(30, userInfo, insuranceSteps)
}
