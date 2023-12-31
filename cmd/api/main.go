package main

import (
	"log"

	"github.com/anggi1792/jec-hospital/domain/hospital"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	dbs "github.com/anggi1792/jec-hospital/pkg/jecconfiguration"
	tools "github.com/anggi1792/jec-hospital/pkg/jectools"
	customvalidator "github.com/anggi1792/jec-hospital/pkg/jecvalidator"
)

func main() {
	//  Load Environtment
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Get Environtment Failed :%v", err)
	}

	//  Database Initial
	dbConn, err := dbs.ConnectSqlx(dbs.DbConfiguration{
		Host:       tools.GetEnv("POSTGRES_HOST"),
		Port:       tools.GetEnv("POSTGRES_PORT"),
		Dbname:     tools.GetEnv("POSTGRES_DBNAME"),
		Dbuser:     tools.GetEnv("POSTGRES_USER"),
		Dbpassword: tools.GetEnv("POSTGRES_PASSWORD"),
		Sslmode:    tools.GetEnv("POSTGRES_SSLMODE"),
	})

	if err != nil {
		panic(err)
	}

	if dbConn == nil {
		panic("Database [" + tools.GetEnv("POSTGRES_DBNAME") + "] Postgree Not Connected!")
	}
	log.Println("Database [" + tools.GetEnv("POSTGRES_DBNAME") + "] Postgree Connected!")

	//  Fiber Framework Initial
	app := fiber.New(
		fiber.Config{
			ErrorHandler: customvalidator.HttpErrorHandler,
		},
	)

	hospital.RouterInitWithDB(app, dbConn)

	log.Println("Hospital API Services Running at port " + tools.GetEnv("BASE_PORT"))
	app.Listen(tools.GetEnv("BASE_PORT"))
}

/*
    ? OUTPUT :
    ? =====================================

{
    "data": {
        var oHospital = hospital.HospitalProto{
		HealthcareId:   "001",
		HealthcareName: "Jec @ Kedoya",
		IsActive:       true,
		UserCreate:     "Anggi",
		CreateAt:       dtNowProto,
	}
    },
    "message": "Create Hospital has been successfully!",
    "status": 201
} */
