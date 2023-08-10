package jecvalidator

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func HttpErrorHandler(c *fiber.Ctx, err error) error {
	report, ok := err.(*fiber.Error)

	if !ok {
		report = fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	var reportMessage interface{} = report.Message

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		var message []string
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				message = append(message, fmt.Sprintf("%s is required",
					err.Field()))
			case "email":
				message = append(message, fmt.Sprintf("%s is not valid email",
					err.Field()))
			case "gte":
				message = append(message, fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param()))
			case "lte":
				message = append(message, fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param()))
			}
		}

		reportMessage = message
	}

	return c.Status(report.Code).JSON(map[string]interface{}{
		"status":  report.Code,
		"message": reportMessage,
	})
}
