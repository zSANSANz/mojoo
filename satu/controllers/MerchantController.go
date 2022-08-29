package controllers

import (
	"net/http"

	"Satu/lib/db"
	"Satu/models"

	"github.com/labstack/echo"
)

func GetMerchantsController(c echo.Context) error {
	merchants, e := db.GetMerchants()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"status":   "success",
		"message":  "success get merchants",
		"merchant": merchants,
	})
}

func GetMerchantByIdController(c echo.Context) error {
	id := c.Param("id")
	merchant, getErr := db.GetMerchantById(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"status":   "success",
		"message":  "success get merchant by id",
		"merchant": merchant,
	})
}

func CreateMerchantController(c echo.Context) error {
	merchant := models.Merchant{}
	c.Bind(&merchant)
	merchants, e := db.CreateMerchant(&merchant)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"status":   "success",
		"message":  "success creating new merchant",
		"merchant": merchants,
	})
}

func UpdateMerchantByIdController(c echo.Context) error {
	id := c.Param("id")
	var UpdateMerchant models.Merchant
	c.Bind(&UpdateMerchant)
	merchant, updateErr := db.UpdateMerchant(id, &UpdateMerchant)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"status":   "success",
		"message":  "success updating existing merchant by id",
		"merchant": merchant,
	})
}

func DeleteMerchantByIdController(c echo.Context) error {
	id := c.Param("id")
	if _, deleteErr := db.DeleteMerchant(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting existing merchant by id",
	})
}
