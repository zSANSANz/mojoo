package controllers

import (
	"net/http"

	"Satu/lib/db"
	"Satu/models"

	"github.com/labstack/echo"
)

func GetOutletsController(c echo.Context) error {
	outlets, e := db.GetOutlets()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get outlets",
		"outlet":  outlets,
	})
}

func GetOutletByIdController(c echo.Context) error {
	id := c.Param("id")
	outlet, getErr := db.GetOutletById(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success get outlet by id",
		"outlet":  outlet,
	})
}

func CreateOutletController(c echo.Context) error {
	outlet := models.Outlet{}
	c.Bind(&outlet)
	outlets, e := db.CreateOutlet(&outlet)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success creating new outlet",
		"outlet":  outlets,
	})
}

func UpdateOutletByIdController(c echo.Context) error {
	id := c.Param("id")
	var UpdateOutlet models.Outlet
	c.Bind(&UpdateOutlet)
	outlet, updateErr := db.UpdateOutlet(id, &UpdateOutlet)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success updating existing outlet by id",
		"outlet":  outlet,
	})
}

func DeleteOutletByIdController(c echo.Context) error {
	id := c.Param("id")
	if _, deleteErr := db.DeleteOutlet(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting existing outlet by id",
	})
}
