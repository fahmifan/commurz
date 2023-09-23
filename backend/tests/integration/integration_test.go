//go:build integration_test

package integration_test

import (
	"database/sql"
	"fmt"
	"net/http"
	"testing"

	"github.com/fahmifan/commurz/pkg/config"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"

	_ "github.com/lib/pq"
)

func TestIntegration(t *testing.T) {
	base := &Base{}
	suite.Run(t, &IntegrationTestSuite{base})
}

type Base struct {
	suite.Suite

	db *sql.DB
}

func (s *Base) SetupSuite() {
	var err error
	config.Parse(".env")
	s.db, err = sql.Open("postgres", config.PostgresDSN())
	s.Require().NoError(err)
}

type IntegrationTestSuite struct {
	*Base
}

func (suite *IntegrationTestSuite) TestRegister() {
	res, err := resty.New().
		SetBaseURL("http://localhost:8080").
		R().
		SetBody(map[string]any{
			"email":           "john@doe.com",
			"name":            "john doe",
			"password":        "test1234",
			"confirmPassword": "test1234",
		}).
		Post("/auth/register")

	suite.NoError(err)
	suite.Equal(http.StatusOK, res.StatusCode())
	fmt.Println(res.String())
}

func (suite *IntegrationTestSuite) TestLogin() {
	res, err := resty.New().
		SetBaseURL("http://localhost:8080").
		R().
		SetBody(map[string]any{
			"email":    "john@doe.com",
			"password": "test1234",
		}).
		Post("/auth")

	suite.NoError(err)
	suite.Equal(http.StatusOK, res.StatusCode())
	fmt.Println(res.String())
}
