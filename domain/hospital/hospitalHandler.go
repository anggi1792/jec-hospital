package hospital

import (
	"github.com/gofiber/fiber/v2"
)

// Declare Hospital Handler construct
type hospitalHandler struct {
	// Access Service
	hospitalSvc hospitalService
}

// Declare Hospital Handler
// @Param appointmentService aka p_oService
func NewHospitalHandler(p_oService hospitalService) hospitalHandler {
	return hospitalHandler{
		hospitalSvc: p_oService,
	}
}

// Create Hospital.
func (h hospitalHandler) Create(ctx *fiber.Ctx) error {
	req := new(HospitalRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}
	if err := ValidateStruct(req); err != nil {
		return ctx.JSON(err)
	}

	err := h.hospitalSvc.CreateService(ctx.Context(), *req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"status":  fiber.StatusCreated,
		"message": "Create Hospital has been successfully!",
		"data":    &req,
	})
}

// Get Hospital By Healthcare ID
func (h hospitalHandler) Get(ctx *fiber.Ctx) error {
	req := new(HospitalRequest)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}
	if err := ValidateStruct(req); err != nil {
		return ctx.JSON(err)
	}

	_, err := h.hospitalSvc.GetService(ctx.Context(), *req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"status":  fiber.StatusOK,
		"message": "Get Hospital has been successfully!",
		"data":    &req,
	})
}

// Get Hospital List.
// func (h hospitalHandler) List(ctx *fiber.Ctx) error {
// 	resp := HospitalProto{}

// 	return ctx.Status(fiber.StatusOK).JSON(map[string]interface{}{
// 		"status":  fiber.StatusOK,
// 		"message": "Get Hospital List has been successfully!",
// 		"data":    &resp,
// 	})
// }

func (h hospitalHandler) List(ctx *fiber.Ctx) error {

	items, err := h.hospitalSvc.ListHospital(ctx.UserContext())
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"status":  fiber.StatusOK,
		"message": "Get Hospital List has been successfully!",
		"data":    items,
	})
}
