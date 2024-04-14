package commons

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// Validate validates the input struct
func Validate(payload interface{}) (*fiber.Error, []*ErrorBody) {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	err := validate.Struct(payload)
	if err != nil {
		// Empty errors slice to store the errors
		var errorList []*ErrorBody
		for _, err := range err.(validator.ValidationErrors) {
			errorList = append(
				errorList,
				&ErrorBody{
					Field:   err.Field(),
					Message: fmt.Sprintf("Invalid data in field %v.", err.Field()),
				},
			)
		}
		return nil, errorList
	}

	return nil, nil
}

// ParseBody is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBody(c *fiber.Ctx, body interface{}) *fiber.Error {
	if err := c.BodyParser(body); err != nil {

		errorList := fiber.ErrBadRequest

		return errorList
	}

	return nil
}

// ParseBodyAndValidate is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBodyAndValidate(c *fiber.Ctx, body interface{}) (*fiber.Error, []*ErrorBody) {

	// First We Parse
	if err := ParseBody(c, body); err != nil {
		return err, nil
	}

	// Then We Validate
	return Validate(body)
}
