package types

type UserInsuranceAnalysis struct {
	Vehicle    string `json:"auto"`
	Disability string `json:"disability"`
	Home       string `json:"home"`
	Life       string `json:"life"`
}
