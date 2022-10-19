package users

import (
	"github.com/coba/dto"
	mock2 "github.com/coba/service/users/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type suiteUsers struct {
	suite.Suite
	handler *HandlerUser
	mock    *mock2.UserMock
}

func (s *suiteUsers) SetupSuite() {
	mock := &mock2.UserMock{}
	s.mock = mock

	s.handler = &HandlerUser{
		UserService: s.mock,
	}
}

func (s *suiteUsers) TestGetAllUsers() {
	s.mock.On("GetAllUser").Return([]dto.DTOUser{
		{
			Name:  "bimo",
			Email: "bimo",
		},
	}, nil)

	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedBody       dto.DTOUser
	}{
		{
			"success",
			http.StatusOK,
			"GET",

			true,
			dto.DTOUser{
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
			r := httptest.NewRequest(v.Method, "/", nil)
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)

			err := s.handler.HandlerUsersAll(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			//if v.HasReturnBody {
			//	var resp map[string]model.Users
			//	err := json.NewDecoder(w.Result().Body).Decode(&resp)
			//
			//	s.NoError(err)
			//	s.Equal(v.ExpectedBody.Name, resp["users"].Name)
			//}
		})
	}
}

func TestSuiteUsers(t *testing.T) {
	suite.Run(t, new(suiteUsers))
}
