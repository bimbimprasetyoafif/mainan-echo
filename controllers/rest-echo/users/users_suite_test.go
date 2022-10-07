package users

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/coba/databases"
	"github.com/coba/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

type suiteUsers struct {
	suite.Suite
	mocking  sqlmock.Sqlmock
	testCase []struct{}
}

func (s *suiteUsers) SetupSuite() {
	dbGormPalsu, mocking, err := sqlmock.New()

	s.NoError(err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      dbGormPalsu,
	}))

	databases.DB = dbGorm
	s.mocking = mocking

}

func (s *suiteUsers) TestGetAllUsers() {
	row := sqlmock.NewRows([]string{"name", "email"}).
		AddRow("bimo ganteng", "bimo@abc.com")

	s.mocking.ExpectQuery(regexp.QuoteMeta("SELECT `users`.`name`,`users`.`email` FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL")).
		WithArgs(1).
		WillReturnRows(row)

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               model.Users
		HasReturnBody      bool
		ExpectedBody       model.Users
	}{
		{
			"success",
			http.StatusOK,
			"POST",
			model.Users{
				Name: "bimo",
			},
			true,
			model.Users{
				Name: "bimo ganteng",
			},
		},
		//{
		//	"bad request",
		//	http.StatusBadRequest,
		//	"GET",
		//	model.Users{},
		//	false,
		//	model.Users{},
		//},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := HandlerUsersByID(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]model.Users
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				s.NoError(err)
				s.Equal(v.ExpectedBody.Name, resp["users"].Name)
			}
		})
	}
}

func (s *suiteUsers) TestDeleteUser() {
	s.mocking.ExpectBegin()

	s.mocking.ExpectExec(regexp.QuoteMeta("DELETE FROM `users` WHERE id = ?")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 0)).
		WillReturnError(nil)

	s.mocking.ExpectCommit()

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedBody       string
	}{
		{
			"success",
			http.StatusOK,
			"DELETE",
			true,
			"not found",
		},
		//{
		//	"bad request",
		//	http.StatusBadRequest,
		//	"GET",
		//	model.Users{},
		//	false,
		//	model.Users{},
		//},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/", nil)
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := HandlerDeleteUsersByID(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]string
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				s.NoError(err)
				s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)
				s.Equal(v.ExpectedBody, resp["message"])
			}
		})
	}
}

func (s *suiteUsers) TearDownSuite() {
	databases.DB = nil
	s.mocking = nil
}

func TestSuiteUsers(t *testing.T) {
	suite.Run(t, new(suiteUsers))
}
