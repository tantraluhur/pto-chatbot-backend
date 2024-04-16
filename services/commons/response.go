package commons

import "github.com/gofiber/fiber/v2"

// Response object as HTTP response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorBody object
type ErrorBody struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// HTTPResponse normalize HTTP Response format
func HTTPResponse(data interface{}) *Response {
	return &Response{
		Success: true,
		Message: "Successfully return",
		Data:    data,
	}
}

// HTTPFiberErrorResponse normalizes error responses
func ValidatorErrorResponse(errorObj []*ErrorBody, message string) *Response {
	// Convert fiber.Error to ErrorBody
	// This fixes issues with swagger auto generated docs not identify fiber.Error type
	var errorSlice []*ErrorBody
	for i := 0; i < len(errorObj); i++ {
		errorSlice = append(errorSlice, mapToErrorOutput(errorObj[i]))
	}

	return &Response{
		Success: false,
		Message: message,
		Data:    errorSlice,
	}
}

func ParserErrorResponse(errorObj *fiber.Error, message string) *Response {

	return &Response{
		Success: false,
		Message: message,
		Data:    errorObj,
	}
}

// HTTPErrorResponse normalizes error responses
func HTTPErrorResponse(message string) *Response {

	return &Response{
		Success: false,
		Message: message,
		Data:    nil,
	}
}

// ==================================== //
// Private Method
func mapToErrorOutput(e *ErrorBody) *ErrorBody {
	return &ErrorBody{
		Field:   e.Field,
		Message: e.Message,
	}
}
