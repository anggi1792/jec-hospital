package hospital

//	Import library
import "time"

//	Declare Hospital Entity
type Hospital struct {
	HealthcareId   string
	HealthcareName string
	IsActive       bool
	UserCreate     string
	CreateAt       time.Time
}
