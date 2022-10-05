package users

import (
	"bytes"
	"encoding/json"
	"github.com/coba/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsersAllHandler(t *testing.T) {
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
		{
			"bad request",
			http.StatusBadRequest,
			"GET",
			model.Users{},
			false,
			model.Users{},
		},
	}

	for _, v := range testCase {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			HandlerUsersAll(w, r)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp model.Users
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				assert.NoError(t, err)

				assert.Equal(t, v.ExpectedBody.Name, resp.Name)
			}
		})
	}
}
