package types

type House struct {
	OwnershipStatus string `json:"ownership_status"`
}

type Vehicle struct {
	Year int `json:"year"`
}

type UserPersonalInformation struct {
	Age           int      `json:"age"`
	Dependents    int      `json:"dependents"`
	House         *House   `json:"house"`
	Income        int      `json:"income"`
	MaritalStatus string   `json:"marital_status"`
	RiskQuestions []int8   `json:"risk_questions"`
	Vehicle       *Vehicle `json:"vehicle"`
}

func (u *UserPersonalInformation) IsUserMarried() bool {
	return u.MaritalStatus == "married"
}

func (u *UserPersonalInformation) IsHouseMortgaged() bool {
	return u.House != nil && u.House.OwnershipStatus == "mortgaged"
}

func (u *UserPersonalInformation) BaseRisk() int {
	baseRisk := 0
	for _, v := range u.RiskQuestions {
		if v != 0 {
			baseRisk++
		}
	}
	return baseRisk
}
