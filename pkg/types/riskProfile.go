package types

type RiskProfile struct {
	Vehicle    string `json:"auto"`
	Disability string `json:"disability"`
	Home       string `json:"home"`
	Life       string `json:"life"`
}

type Step struct {
	Risk        int
	Eligibility bool
}

type UserInsuranceAnalysisSteps struct {
	Auto       Step
	Disability Step
	Home       Step
	Life       Step
}
