package main

import (
	"log"
	"net"

	"github.com/anggi1792/jec-hospital/domain/hospital"
	dbs "github.com/anggi1792/jec-hospital/pkg/jecconfiguration"
	tools "github.com/anggi1792/jec-hospital/pkg/jectools"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	//  Load Environtment
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Get Environtment Failed :%v", err)
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

	srv := grpc.NewServer()
	hospital.RouterInitGRPC(srv, dbConn)

	log.Println("Registered GRPC Route ...")
	listen, err := net.Listen("tcp", tools.GetEnv("GRPC_PORT"))
	if err != nil {
		panic(err)
	}

	log.Println("Hospital GRPC Server Running at port ", tools.GetEnv("GRPC_PORT"))
	err = srv.Serve(listen)
	if err != nil {
		panic(err)
	}
}
