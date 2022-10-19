package users

import (
	"fmt"
	"github.com/coba/databases"
	"github.com/coba/model"
	"github.com/coba/service/users"
	"github.com/labstack/echo/v4"
	"strconv"
)

type DTOUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type HandlerUser struct {
	UserService users.UserService
}

func (h *HandlerUser) HandlerUsersAll(c echo.Context) error {
	u, err := h.UserService.GetAllUser()
	if err != nil {
		fmt.Println(err.Error())

		return c.JSON(500, map[string]interface{}{
			"message": "something wrong",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"users": u,
	})
}

func HandlerCreateUsers(c echo.Context) error {
	var u DTOUser

	_ = c.Bind(&u)

	databases.DB.Create(&model.Users{
		Name:  u.Name,
		Email: u.Email,
	})

	return c.JSON(200, map[string]interface{}{
		"message": "created",
	})
}

func HandlerUpdateUsers(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var u DTOUser

	err = c.Bind(&u)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	databases.DB.Model(&model.Users{}).
		Where("id = ?", idInt).
		Updates(&model.Users{
			Name:  u.Name,
			Email: u.Email})

	return c.JSON(200, map[string]interface{}{
		"message": "updated",
	})
}

func HandlerUsersByID(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	temp := DTOUser{}

	databases.DB.Model(&model.Users{}).Where("id = ?", idInt).Find(&temp)

	return c.JSON(200, map[string]interface{}{
		"users": temp,
	})
}

func HandlerDeleteUsersByID(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	res := databases.DB.Unscoped().Delete(&model.Users{}, "id = ?", idInt)
	if res.Error != nil {
		return c.JSON(500, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if res.RowsAffected < 1 {
		return c.JSON(200, map[string]interface{}{
			"message": "not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "success",
	})
}
