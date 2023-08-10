package hospital

import (
	context "context"
	"log"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//	Import library

// Declare Hospitall Handler construct
type hospitalGrpcHandler struct {
	// Access Service
	hospitalSvc hospitalService
}

// Declare Hospital GRPC Handler
// @Param hospitalService
func NewHospitalGrpcHandler(p_oGrpcService hospitalService) hospitalGrpcHandler {
	return hospitalGrpcHandler{
		hospitalSvc: p_oGrpcService,
	}
}

// ADD
func (gh *hospitalGrpcHandler) Add(ctx context.Context, req *HospitalAddProto) (res *HospitalProto, err error) {
	reqData := HospitalRequest{
		HealthcareId:   req.Addhospital.HealthcareId,
		HealthcareName: req.Addhospital.HealthcareName,
		IsActive:       req.Addhospital.IsActive,
		UserCreate:     req.Addhospital.UserCreate,
		CreateAt:       req.Addhospital.CreateAt.AsTime(),
	}

	err = gh.hospitalSvc.CreateService(ctx, reqData)
	res = req.Addhospital

	return
}

// GET
func (gh *hospitalGrpcHandler) Get(ctx context.Context, filter *HospitalGetProto) (res *HospitalProto, err error) {
	getData := HospitalRequest{
		HealthcareId: filter.GetHealthcareId(),
	}

	dataResult, err := gh.hospitalSvc.GetService(ctx, getData)
	if len(dataResult) == 0 {
		return
	}

	log.Printf("cek %+v test", dataResult)

	res = &HospitalProto{
		HealthcareId:   dataResult[0].HealthcareId,
		HealthcareName: dataResult[0].HealthcareName,
		IsActive:       dataResult[0].IsActive,
		UserCreate:     dataResult[0].UserCreate,
		CreateAt:       timestamppb.New(dataResult[0].CreateAt),
	}
	return
}

// LIST
// func (gh *hospitalGrpcHandler) List(ctx context.Context, empty *emptypb.Empty) (res *HospitalListProto, err error) {
// 	return
// }

func (gh *hospitalGrpcHandler) List(ctx context.Context, empty *emptypb.Empty) (res *HospitalListProto, err error) {
	// getData := HospitalRequest{
	// 	HealthcareId: list.GetHealthcareId(),
	// }

	dataResult, err := gh.hospitalSvc.ListHospital(ctx)
	if len(dataResult) == 0 {
		return
	}

	var resultdata []*HospitalProto

	for _, v := range dataResult {
		resultdata = append(resultdata, &HospitalProto{
			HealthcareId:   v.HealthcareId,
			HealthcareName: v.HealthcareName,
			IsActive:       v.IsActive,
			UserCreate:     v.UserCreate,
			CreateAt:       timestamppb.New(v.CreateAt),
		})
	}

	log.Printf("cek %+v test", dataResult)

	res = &HospitalListProto{
		Hospitals: resultdata,
	}
	return
}

func (gh *hospitalGrpcHandler) mustEmbedUnimplementedHospitalServiceServer() {}
