package hospital

//	Import library
import "time"

//	Declare Hospital Entity
type Hospital struct {
	HealthcareId   string    `db:"healthcare_id"`
	HealthcareName string    `db:"healthcare_name"`
	IsActive       bool      `db:"is_active"`
	UserCreate     string    `db:"user_create"`
	CreateAt       time.Time `db:"create_at"`
}
