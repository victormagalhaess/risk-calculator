package riskSteps

import (
	"time"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/types"
)

func NoVehicle(userInfo types.UserPersonalInformation,
	insuranceSteps *types.UserInsuranceAnalysisSteps) {
	if userInfo.Vehicle == nil {
		insuranceSteps.Auto.Eligibility = false
	}
}

func VehicleBuiltLast5Years(userInfo types.UserPersonalInformation,
	insuranceSteps *types.UserInsuranceAnalysisSteps) {
	currentYear, _, _ := time.Now().Date()
	yearDifference := currentYear - userInfo.Vehicle.Year
	if userInfo.Vehicle != nil && yearDifference < 5 {
		insuranceSteps.Auto.Risk++
	}
}
