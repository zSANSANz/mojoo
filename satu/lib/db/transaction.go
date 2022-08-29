package db

import (
	"Satu/config"
	"Satu/models"
)

func GetTransactions() (interface{}, error) {
	var transaction []models.Transaction

	if err := config.DB.Find(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
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
