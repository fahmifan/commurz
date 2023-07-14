package webserver

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	commurzpbv1 "github.com/fahmifan/commurz/pkg/pb/commurz/v1"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	*Webserver
}

func (handler HomeHandler) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "home/index.html", echo.Map{
			"name": "john doe",
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
			return err
		}

		listRes, err := handler.service.ListUsers(
			c.Request().Context(),
			&connect.Request[commurzpbv1.ListUsersRequest]{},
		)
		if err != nil {
			return err
		}

		return c.Render(http.StatusOK, "user/created.html", echo.Map{
			"users": listRes.Msg.Users,
		})
	}
}
