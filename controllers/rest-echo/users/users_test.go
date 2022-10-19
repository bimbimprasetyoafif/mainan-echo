package users

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/coba/databases"
	"github.com/coba/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestUsersAllHandler(t *testing.T) {
	// mocking
	t.Log()
	dbGormPalsu, mocking, err := sqlmock.New()

	assert.NoError(t, err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      dbGormPalsu,
	}))

	databases.DB = dbGorm

	row := sqlmock.NewRows([]string{"name", "email"}).
		AddRow("bimo ganteng", "bimo@abc.com")

	mocking.ExpectQuery(regexp.QuoteMeta("SELECT `users`.`name`,`users`.`email` FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL")).
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
		t.Run(v.Name, func(t *testing.T) {
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
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]model.Users
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				assert.NoError(t, err)
				assert.Equal(t, v.ExpectedBody.Name, resp["users"].Name)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	// mocking
	dbGormPalsu, mocking, err := sqlmock.New()

	assert.NoError(t, err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      dbGormPalsu,
	}))

	databases.DB = dbGorm

	mocking.ExpectBegin()

	mocking.ExpectExec(regexp.QuoteMeta("DELETE FROM `users` WHERE id = ?")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 0)).
		WillReturnError(nil)

	mocking.ExpectCommit()

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
		t.Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/", nil)
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := HandlerDeleteUsersByID(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]string
				err := json.NewDecoder(w.Result().Body).Decode(&resp)

				assert.NoError(t, err)
				assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)
				assert.Equal(t, v.ExpectedBody, resp["message"])
			}
		})
	}
}

//
//type DBPalsu struct {
//	mock.Mock
//}
//
//func (d *DBPalsu) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
//	panic("")
//}
//func (d *DBPalsu) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
//	panic("")
//}
//func (d *DBPalsu) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
//	panic("")
//}
//func (d *DBPalsu) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
//	call := d.Called(ctx, query, args)
//
//	return call.Get(0).(*sql.Row)
//}
//
//type DBPalsuError struct {
//	FuncQueryContext func(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
//}
//
//func (d *DBPalsuError) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
//	panic("")
//}
//func (d *DBPalsuError) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
//	panic("")
//}
//func (d *DBPalsuError) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
//	return d.FuncQueryContext(ctx, query, args...)
//}
//func (d *DBPalsuError) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
//	panic("")
//}
