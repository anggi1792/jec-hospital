package main

import (
	"log"
	"net"

	"github.com/anggi1792/jec-hospital/domain/hospital"
	"github.com/anggi1792/jec-hospital/domain/paramedics"
	"github.com/anggi1792/jec-hospital/domain/schedules"
	dbs "github.com/anggi1792/jec-hospital/pkg/jecconfiguration"
	tools "github.com/anggi1792/jec-hospital/pkg/jectools"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	//  Load Environtment
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Get Environtment Failed :%v", err)
	}

	//  Database Initial Hospital
	dbConnHospital, err := dbs.ConnectSqlx(dbs.DbConfiguration{
		Host:       tools.GetEnv("POSTGRES_HOST"),
		Port:       tools.GetEnv("POSTGRES_PORT"),
		Dbname:     tools.GetEnv("POSTGRES_DBNAME_Hospital"),
		Dbuser:     tools.GetEnv("POSTGRES_USER"),
		Dbpassword: tools.GetEnv("POSTGRES_PASSWORD"),
		Sslmode:    tools.GetEnv("POSTGRES_SSLMODE"),
	})

	if err != nil {
		panic(err)
	}

	// PANGGIL DB HOSPITAL
	if dbConnHospital == nil {
		panic("Database [" + tools.GetEnv("POSTGRES_DBNAME_Hospital") + "] Postgree Not Connected!")
	}
	log.Println("Database [" + tools.GetEnv("POSTGRES_DBNAME_Hospital") + "] Postgree Connected!")

	//  Database Initial Paramedic
	dbConnParamedics, err := dbs.ConnectSqlx(dbs.DbConfiguration{
		Host:       tools.GetEnv("POSTGRES_HOST"),
		Port:       tools.GetEnv("POSTGRES_PORT"),
		Dbname:     tools.GetEnv("POSTGRES_DBNAME_PARAMEDICS"),
		Dbuser:     tools.GetEnv("POSTGRES_USER"),
		Dbpassword: tools.GetEnv("POSTGRES_PASSWORD"),
		Sslmode:    tools.GetEnv("POSTGRES_SSLMODE"),
	})

	if err != nil {
		panic(err)
	}

	// PANGGIL DB PARAMEDICS
	if dbConnParamedics == nil {
		panic("Database [" + tools.GetEnv("POSTGRES_DBNAME_PARAMEDICS") + "] Postgree Not Connected!")
	}
	log.Println("Database [" + tools.GetEnv("POSTGRES_DBNAME_PARAMEDICS") + "] Postgree Connected!")

	srv := grpc.NewServer()
	hospital.RouterInitGRPC(srv, dbConnHospital)
	paramedics.RegisterRouteGRPC(srv, dbConnParamedics)
	schedules.RegisterRouteGRPC(srv, dbConnParamedics)

	log.Println("Registered GRPC Route ...")
	listen, err := net.Listen("tcp", tools.GetEnv("GRPC_PORT"))
	if err != nil {
		panic(err)
	}

	log.Println("Hospital & Paramedics GRPC Server Running at port ", tools.GetEnv("GRPC_PORT"))
	err = srv.Serve(listen)
	if err != nil {
		panic(err)
	}
}
