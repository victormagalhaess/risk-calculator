package types

type house struct {
	OwnershipStatus string `json:"ownership_status"`
}

type vehicle struct {
	Year int `json:"year"`
}

type UserPersonalInformation struct {
	Age           int     `json:"age"`
	Dependents    int     `json:"dependents"`
	House         house   `json:"house"`
	Income        int     `json:"income"`
	MaritalStatus string  `json:"marital_status"`
	RiskQuestions []bool  `json:"risk_questions"`
	Vehicle       vehicle `json:"vehicle"`
}

func (u *UserPersonalInformation) IsUserMarried() bool {
	return u.MaritalStatus == "married"
}

func (u *UserPersonalInformation) IsHouseMortgaged() bool {
	return u.House.OwnershipStatus == "mortgaged"
}

func (u *UserPersonalInformation) BaseRisk() int {
	baseRisk := 0
	for _, v := range u.RiskQuestions {
		if v {
			baseRisk++
		}
	}
	return baseRisk
}
