package hospital

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

func RouterInit(r fiber.Router) {
	svc := NewHospitalService()
	handler := NewHospitalHandler(svc)

	hospitalApi := r.Group("/api")
	hospitalApi.Get("/hospital", handler.List)
	hospitalApi.Get("/hospital", handler.Get)
	hospitalApi.Post("/hospital", handler.Create)
}

func RouterInitWithDB(r fiber.Router, dbx *sqlx.DB) {
	repo := NewHospitalRepository(dbx)
	svc := NewHospitalServiceDB(repo)
	handler := NewHospitalHandler(svc)

	hospitalApi := r.Group("/api")
	hospitalApi.Get("/hospital", handler.List)
	hospitalApi.Get("/hospital", handler.Get)
	hospitalApi.Post("/hospital", handler.Create)
}

func RouterInitGRPC(server *grpc.Server, db *sqlx.DB) {
	repo := NewHospitalRepository(db)
	svc := NewHospitalServiceDB(repo)
	handler := NewHospitalGrpcHandler(svc)

	RegisterHospitalServiceServer(server, &handler)
}
