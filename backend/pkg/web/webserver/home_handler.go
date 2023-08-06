package webserver

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/fahmifan/commurz/pkg/logs"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	*Webserver
}

func (handler HomeHandler) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		listUserRes, err := handler.service.ListUsers(
			c.Request().Context(),
			&connect.Request[commurzpbv1.ListUsersRequest]{},
		)
		if err != nil {
			// TODO: log error
			return c.Render(http.StatusInternalServerError, "partials/error_toast.html", echo.Map{
				"message": "something went wrong",
			})
		}

		return c.Render(http.StatusOK, "home/index.html", echo.Map{
			"users": listUserRes.Msg.Users,
		})
	}
}

type UserHandler struct {
	*Webserver
}

func (handler UserHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &commurzpbv1.CreateUserRequest{
			Email: c.FormValue("email"),
		}

		_, err := handler.service.CreateUser(
			c.Request().Context(),
			&connect.Request[commurzpbv1.CreateUserRequest]{Msg: req},
		)
		if err != nil {
			logs.ErrCtx(c.Request().Context(), err, "UserHandler-Create-CreateUser")
			return c.Render(http.StatusBadRequest, "home/index.html", echo.Map{
				"message": "something went wrong",
			})
		}

		listRes, err := handler.service.ListUsers(
			c.Request().Context(),
			&connect.Request[commurzpbv1.ListUsersRequest]{},
		)
		if err != nil {
			logs.ErrCtx(c.Request().Context(), err, "UserHandler-Create-ListUsers")
			return c.Render(http.StatusInternalServerError, "home/index.html", echo.Map{
				"message": "something went wrong",
			})
		}

		return c.Render(http.StatusOK, "home/index.html", echo.Map{
			"users": listRes.Msg.Users,
		})
	}
}
