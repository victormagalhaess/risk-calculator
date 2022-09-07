package model

// @Description Risk profile for a user
type RiskProfile struct {
	Vehicle    string `json:"auto" example:"economic"`
	Disability string `json:"disability" example:"regular"`
	Home       string `json:"home,omitempty" example:"ineligible"`
	Renters    string `json:"renters,omitempty" example:"ineligible"`
	Life       string `json:"life" example:"responsible"`
}

type StepResult struct {
	Risk        int
	Eligibility bool
}

type UserInsuranceAnalysisSteps struct {
	Auto       StepResult
	Disability StepResult
	Home       StepResult
	Life       StepResult
}

func mapStepToRiskProfile(step StepResult) string {
	if !step.Eligibility {
		return "ineligible"
	}
	if step.Risk <= 0 {
		return "economic"
	}
	if step.Risk <= 2 {
		return "regular"
	}
	return "responsible"
}

func mapHouseToRiskProfile(home StepResult, house *House, riskProfile RiskProfile) RiskProfile {
	if house != nil && house.OwnershipStatus == "rented" {
		riskProfile.Renters = mapStepToRiskProfile(home)
		return riskProfile
	}
	riskProfile.Home = mapStepToRiskProfile(home)
	return riskProfile

}

func (u *UserInsuranceAnalysisSteps) MapInsuranceAnalisysToRiskProfile(userInfo UserPersonalInformation) RiskProfile {
	riskProfile := RiskProfile{}
	riskProfile.Vehicle = mapStepToRiskProfile(u.Auto)
	riskProfile.Disability = mapStepToRiskProfile(u.Disability)
	riskProfile.Life = mapStepToRiskProfile(u.Life)
	return mapHouseToRiskProfile(u.Home, userInfo.House, riskProfile)
}

func (u *UserInsuranceAnalysisSteps) SetEligibility(eligibility bool) {
	u.Auto.Eligibility = eligibility
	u.Disability.Eligibility = eligibility
	u.Home.Eligibility = eligibility
	u.Life.Eligibility = eligibility
}
