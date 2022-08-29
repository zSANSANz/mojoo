package controllers

import (
	"net/http"

	"Satu/lib/db"
	"Satu/models"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	users, e := db.GetUsers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get users",
		"user":    users,
	})
}

func GetUserByIdController(c echo.Context) error {
	id := c.Param("id")
	user, getErr := db.GetUserById(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get user by id",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	users, e := db.CreateUser(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success creating new user",
		"user":    users,
	})
}

func UpdateUserByIdController(c echo.Context) error {
	id := c.Param("id")
	var UpdateUser models.User
	c.Bind(&UpdateUser)
	user, updateErr := db.UpdateUser(id, &UpdateUser)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success updating existing user by id",
		"user":    user,
	})
}

func DeleteUserByIdController(c echo.Context) error {
	id := c.Param("id")
	if _, deleteErr := db.DeleteUser(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting existing user by id",
	})
}

func RegisterUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	_, err := db.RegisterUser(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success registring new user",
		"user":    user,
	})
}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	users, e := db.LoginUser(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":   200,
		"status": "success login",
		"users":  users,
	})
}
