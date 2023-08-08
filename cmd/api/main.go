package main

import (
	"log"
	"time"

	"github.com/anggi1792/jec-hospital/domain/hospital"
	"github.com/golang/protobuf/ptypes"
)

func main() {
	dtDateTimeNow := time.Now()

	dtNowProto, err := ptypes.TimestampProto(dtDateTimeNow)

	if err != nil {
		log.Println(err)
	}

	var oHospital = hospital.HospitalProto{
		HealthcareId:   "001",
		HealthcareName: "Jec @ Kedoya",
		IsActive:       true,
		UserCreate:     "Anggi",
		CreateAt:       dtNowProto,
	}

	var oHosptalList = hospital.HospitalListProto{
		Hospitals: []*hospital.HospitalProto{&oHospital},
	}

	log.Printf("data: \n%v", oHosptalList.Hospitals)
}
