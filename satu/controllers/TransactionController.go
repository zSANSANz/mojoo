package controllers

import (
	"net/http"

	"Satu/lib/db"
	"Satu/models"

	"github.com/labstack/echo"
)

func GetTransactionsController(c echo.Context) error {
	transactions, e := db.GetTransactions()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":        200,
		"status":      "success",
		"message":     "success get transactions",
		"transaction": transactions,
	})
}

func GetTransactionByIdController(c echo.Context) error {
	id := c.Param("id")
	transaction, getErr := db.GetTransactionById(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":        200,
		"status":      "success",
		"message":     "success get transaction by id",
		"transaction": transaction,
	})
}

func CreateTransactionController(c echo.Context) error {
	transaction := models.Transaction{}
	c.Bind(&transaction)
	transactions, e := db.CreateTransaction(&transaction)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":        200,
		"status":      "success",
		"message":     "success creating new transaction",
		"transaction": transactions,
	})
}

func UpdateTransactionByIdController(c echo.Context) error {
	id := c.Param("id")
	var UpdateTransaction models.Transaction
	c.Bind(&UpdateTransaction)
	transaction, updateErr := db.UpdateTransaction(id, &UpdateTransaction)
	if updateErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, updateErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":        200,
		"status":      "success",
		"message":     "success updating existing transaction by id",
		"transaction": transaction,
	})
}

func DeleteTransactionByIdController(c echo.Context) error {
	id := c.Param("id")
	if _, deleteErr := db.DeleteTransaction(id); deleteErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, deleteErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "success deleting existing transaction by id",
	})
}
