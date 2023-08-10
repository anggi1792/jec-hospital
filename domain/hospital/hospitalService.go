package hospital

import (
	"context"
	"log"
)

type Repository interface {
	AddRepo(ctx context.Context, h Hospital) (p_strHealthcareId string, err error)
	GetRepo(ctx context.Context, healthcareid string) (res []Hospital, err error)
	GetListHospital(ctx context.Context) ([]Hospital, error)
}

// Declare Hospital Service construct
type hospitalService struct {
	// Access ke repo
	repo Repository
}

// Declare Hospital Service
func NewHospitalService() hospitalService {
	return hospitalService{}
}

func NewHospitalServiceDB(oRepo Repository) hospitalService {
	return hospitalService{
		repo: oRepo,
	}
}

// #region Business Logic Method

// Implement Create Hospital on handler
func (asvc hospitalService) CreateService(ctx context.Context, req HospitalRequest) (err error) {
	model := req.ParseToEntity()

	healthcareid, err := asvc.repo.AddRepo(ctx, model)
	if err != nil {
		return
	}

	log.Printf("Create Hospital [%v] successfully!\n", healthcareid)

	return nil
}

// Implement Get Hospital on handler
func (asvc hospitalService) GetService(ctx context.Context, req HospitalRequest) (res []Hospital, err error) {
	model := req.ParseToEntity()

	res, err = asvc.repo.GetRepo(ctx, model.HealthcareId)
	if err != nil {
		return
	}

	log.Printf("Create Hospital [%v] successfully!\n", res)

	return
}

// Implement List Hospital on handler
// func (asvc hospitalService) ListHospital(ctx context.Context, req HospitalRequest) (res []Hospital, err error) {
// 	model := req.ParseToEntity()

// 	res, err = asvc.repo.GetListHospital(ctx, model.HealthcareId)
// 	if err != nil {
// 		return
// 	}

// 	log.Printf("Create Hospital [%v] successfully!\n", res)

// 	return
// }

func (asvc hospitalService) ListHospital(ctx context.Context) ([]HospitalListResponse, error) {
	result := make([]HospitalListResponse, 0)

	items, err := asvc.repo.GetListHospital(ctx)
	if err != nil {
		return result, err
	}

	result = make([]HospitalListResponse, len(items))
	for i, item := range items {
		result[i].HealthcareId = item.HealthcareId
		result[i].HealthcareName = item.HealthcareName
		result[i].IsActive = item.IsActive
		result[i].UserCreate = item.UserCreate
		result[i].CreateAt = item.CreateAt
	}

	return result, nil
}

// #endregion end of Business Logic Method
