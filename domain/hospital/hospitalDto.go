package hospital

//	Import library
import (
	"time"
)

// [Response] Declare Hospital List construct
type HospitalListResponse struct {
	HealthcareId   string    `json:"HealthcareId"`
	HealthcareName string    `json:"HealthcareName"`
	IsActive       bool      `json:"IsActive"`
	UserCreate     string    `json:"UserCreate"`
	CreateAt       time.Time `json:"CreateAt"`
}

// [Request] Declare Hospital List construct
type HospitalRequest struct {
	HealthcareId   string    `json:"HealthcareId"`
	HealthcareName string    `json:"HealthcareName"`
	IsActive       bool      `json:"IsActive"`
	UserCreate     string    `json:"UserCreate"`
	CreateAt       time.Time `json:"CreateAt"`
}

//  #region Method

// Convert Object to Entity
func (a HospitalRequest) ParseToEntity() Hospital {
	return Hospital{
		HealthcareId:   a.HealthcareId,
		HealthcareName: a.HealthcareName,
		IsActive:       a.IsActive,
		UserCreate:     a.UserCreate,
		CreateAt:       time.Now(),
	}
}
