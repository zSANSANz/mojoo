package db

import (
	"Satu/config"
	"Satu/models"
	"fmt"
	"math"
	"net/http"
	"strings"
)

func GetTransactions() (interface{}, error) {
	var transaction []models.Transaction

	if err := config.DB.Find(&transaction).Error; err != nil {
		return nil, err
	}

	paginatedData, _ := paginate.New(transaction).Find()

	return paginatedData, nil
}

func GetTransactionById(id string) (interface{}, error) {
	var transaction []models.Transaction

	if err := config.DB.First(&transaction, "transaction_id=?", id).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func CreateTransaction(transaction *models.Transaction) (interface{}, error) {
	if err := config.DB.Save(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func UpdateTransaction(id string, transaction *models.Transaction) (interface{}, error) {
	var existingTransaction models.Transaction
	if err := config.DB.First(&existingTransaction, "transaction_id=?", id).Error; err != nil {
		return nil, err
	}
	existingTransaction.MerchantId = transaction.MerchantId
	existingTransaction.OutletId = transaction.OutletId
	existingTransaction.BillTotal = transaction.BillTotal
	if updateErr := config.DB.Save(&existingTransaction).Where("transaction_id=?", id).Error; updateErr != nil {
		return nil, updateErr
	}
	return existingTransaction, nil
}

func DeleteTransaction(id string) (interface{}, error) {
	var transaction models.Transaction
	if err := config.DB.First(&transaction, "transaction_id=?", id).Error; err != nil {
		return nil, err
	}
	if deleteErr := config.DB.Delete(&transaction).Where("transaction_id=?", id).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return transaction, nil
}

func paginate(value interface{}, pagination *models.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

type TransactionGorm struct {
	db *gorm.DB
}

func (cg *TransactionGorm) List(pagination pkg.Pagination) (*pkg.Pagination, error) {
	var transaction []*Transaction

	tokenString := ""

	// get jwt bearer token from context
	for key, values := range c.Request().Header {
		for _, value := range values {
			if key == "Authorization" {
				tokenString = strings.Replace(value, "Bearer ", "", -1)

			}
		}
	}

	// decode token with secret code bermaslaah
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("bermaslaah"), nil
	})

	if !token.Valid {
		fmt.Println("invalid jwt token!!!")
	}

	if users.id != transaction.id {
		// fmt.Println("kamu bukanlah head")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":     "false",
			"message":    "Unauthorized: Access is denied due to invalid credentials",
			"error_code": "401",
			"data":       users,
		})
	} else {
		cg.db.Scopes(paginate(transaction, &pagination, cg.db)).Find(&transaction)
		pagination.Rows = transaction
	}

	return &pagination, nil
}
