package types

type RiskProfile struct {
	Vehicle    string `json:"auto"`
	Disability string `json:"disability"`
	Home       string `json:"home"`
	Life       string `json:"life"`
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
	if step.Eligibility == false {
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

func (u *UserInsuranceAnalysisSteps) MapInsuranceAnalisysToRiskProfile() RiskProfile {
	riskProfile := RiskProfile{}
	riskProfile.Vehicle = mapStepToRiskProfile(u.Auto)
	riskProfile.Disability = mapStepToRiskProfile(u.Disability)
	riskProfile.Home = mapStepToRiskProfile(u.Home)
	riskProfile.Life = mapStepToRiskProfile(u.Life)
	return riskProfile
}
