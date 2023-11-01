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
