package helper

import (
	"github.com/labstack/echo"
)

type (

	// Response - structure response
	Response struct {
		Ctx     echo.Context           `json:"-"`
		Status  int                    `json:"status"`
		Code    int                    `json:"code"`
		Message string                 `json:"message"`
		Data    map[string]interface{} `json:"data,omitempty"`
	}
)

// NewResponse - instantiate new Response
// status: http status code
// code: internal error code dictionary
func NewResponse(status, code int, msg string, data map[string]interface{}) *Response {
	return &Response{
		Status:  status,
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

// WriteResponse - write response to the client
func (r *Response) WriteResponse(ctx echo.Context) error {
	return ctx.JSON(r.Status, r)
}
