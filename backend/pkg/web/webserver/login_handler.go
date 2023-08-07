package webserver

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/fahmifan/commurz/pkg/internal/pkguser"
	"github.com/fahmifan/commurz/pkg/logs"
	"github.com/fahmifan/commurz/pkg/web/webserver/httperr"
	"github.com/labstack/echo/v4"
)

type Login struct {
	*Webserver
}

func (handler *Login) Create() echo.HandlerFunc {
	return func(ec echo.Context) error {
		req := struct {
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required"`
		}{}
		if err := ec.Bind(&req); err != nil {
			return ec.JSON(http.StatusBadRequest, httperr.ErrInvalidArg)
		}

		userReader := pkguser.UserReader{}

		user := pkguser.User{}
		user, err := userReader.FindByEmail(ec.Request().Context(), handler.db, req.Email)
		if err != nil {
			if isNotFoundErr(err) {
				return ec.JSON(http.StatusNotFound, httperr.ErrNotFound)
			}

			logs.ErrCtx(ec.Request().Context(), err, "Login: Create: userReader.FindByEmail")
			return ec.JSON(http.StatusBadRequest, httperr.ErrInternal)
		}

		err = user.CanLogin(req.Password)
		if err != nil {
			return ec.JSON(http.StatusBadRequest, httperr.ErrInvalidCredential)
		}

		err = handler.session.SaveUser(ec, _sessionUserMaxAge, user.ID)
		if err != nil {
			logs.ErrCtx(ec.Request().Context(), err, "Login: Create: session.SaveUser")
			return ec.JSON(http.StatusBadRequest, httperr.ErrInternal)
		}

		return ec.JSON(http.StatusOK, echo.Map{
			"message": "success",
		})
	}
}

func (a *Login) Delete() echo.HandlerFunc {
	return func(ec echo.Context) error {
		err := a.session.DeleteUser(ec)
		if err != nil {
			return ec.Redirect(http.StatusFound, ec.Echo().Reverse("page-auth-login-new"))
		}

		return ec.Redirect(http.StatusFound, "/")
	}
}

func isNotFoundErr(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
