package hospital

import (
	context "context"

	tools "github.com/anggi1792/jec-hospital/pkg/jectools"
	"github.com/jmoiron/sqlx"
)

type hospitalRepository struct {
	db *sqlx.DB
}

// Declare Hospital Repository
func NewHospitalRepository(p_oDatabase *sqlx.DB) hospitalRepository {
	return hospitalRepository{db: p_oDatabase}
}

// Implements Add Hospital Repository on Service
func (r hospitalRepository) AddRepo(ctx context.Context, app Hospital) (p_strHealthcareId string, err error) {
	query := `INSERT INTO healthcares (
        healthcare_id, healthcare_name, is_active, user_create, create_at)
        VALUES ($1, $2, $3, $4, $5) returning healthcare_id`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return "-99", err
	}

	err = stmt.QueryRow(
		tools.NewSQLNullString(app.HealthcareId),
		tools.NewSQLNullString(app.HealthcareName),
		app.IsActive,
		tools.NewSQLNullString(app.UserCreate),
		app.CreateAt,
	).Scan(&p_strHealthcareId)

	if err != nil {
		return "-99", err
	}

	return
}

// Implements GET Hospital Repository on Service

func (r hospitalRepository) GetRepo(ctx context.Context, healthcareid string) ([]Hospital, error) {
	var (
		res  []Hospital
		args []any
	)

	query := `SELECT * FROM healthcares WHERE healthcare_id = ?`
	args = append(args, healthcareid)

	if healthcareid != "" {
		query += ` AND healthcare_id = ?`
		args = append(args, healthcareid)
	}

	query = r.db.Rebind(query)

	if err := r.db.SelectContext(ctx, &res, query, args...); err != nil {
		return res, err
	}

	return res, nil
}

// Implements LIST Hospital Repository on Service
func (r hospitalRepository) GetListHospital(ctx context.Context) ([]Hospital, error) {
	var (
		result = make([]Hospital, 0)
	)

	query := `SELECT * FROM healthcares `
	query = r.db.Rebind(query)

	if err := r.db.SelectContext(ctx, &result, query); err != nil {
		return result, err
	}

	return result, nil
}
