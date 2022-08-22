package types

type House struct {
	OwnershipStatus string `json:"ownership_status"`
}

type Vehicle struct {
	Year int `json:"year"`
}

// @Description User Personal information
type UserPersonalInformation struct {
	Age           int      `json:"age" example:"30"`
	Dependents    int      `json:"dependents" example:"2"`
	House         *House   `json:"house" swaggertype:"object,string" example:"ownership_status:mortgaged"`
	Income        int      `json:"income" example:"100000"`
	MaritalStatus string   `json:"marital_status" example:"married"`
	RiskQuestions []int8   `json:"risk_questions" swaggertype:"array,integer"  example:"1,0,0"`
	Vehicle       *Vehicle `json:"vehicle" swaggertype:"object,integer" example:"year:2018"`
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
