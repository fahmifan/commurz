package httperr

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Errs struct {
	Msg        string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Code       int    `json:"code"`
}

func (errs Errs) WithMessage(msg string) Errs {
	errs.Msg = msg
	return errs
}

func WithMessage(err error, msg string) error {
	errs, ok := err.(Errs)
	if !ok {
		return err
	}

	errs.Msg = msg
	return errs
}

func (e Errs) Error() string {
	return fmt.Sprintf("status-code:%d message:%s", e.StatusCode, e.Msg)
}

func JSON(ec echo.Context, errs Errs) error {
	return ec.JSON(errs.StatusCode, errs)
}

var (
	ErrNotFound          = Errs{Code: 1001, StatusCode: http.StatusNotFound, Msg: "Not Found"}
	ErrInvalidArg        = Errs{Code: 1002, StatusCode: http.StatusBadRequest, Msg: "invalid argument"}
	ErrInternal          = Errs{Code: 1003, StatusCode: http.StatusInternalServerError, Msg: "internal error"}
	ErrUnauthorized      = Errs{Code: 1004, StatusCode: http.StatusUnauthorized, Msg: "unauthorized"}
	ErrInvalidCredential = Errs{Code: 1005, StatusCode: http.StatusUnauthorized, Msg: "invalid credential"}
)
