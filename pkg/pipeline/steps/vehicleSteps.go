package steps

import (
	"time"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"
)

func NoVehicle(userInfo model.UserPersonalInformation,
	insuranceSteps *model.UserInsuranceAnalysisSteps) {
	if userInfo.Vehicle == nil {
		insuranceSteps.Auto.Eligibility = false
	}
}

func VehicleBuiltLast5Years(userInfo model.UserPersonalInformation,
	insuranceSteps *model.UserInsuranceAnalysisSteps) {
	currentYear, _, _ := time.Now().Date()
	if userInfo.Vehicle != nil && currentYear-userInfo.Vehicle.Year < 5 {
		insuranceSteps.Auto.Risk++
	}
}
