package db

import (
	"Satu/config"
	"Satu/models"
)

func GetMerchants() (interface{}, error) {
	var merchant []models.Merchant

	if err := config.DB.Find(&merchant).Error; err != nil {
		return nil, err
	}
	return merchant, nil
}

func GetMerchantById(id string) (interface{}, error) {
	var merchant []models.Merchant

	if err := config.DB.First(&merchant, "merchant_id=?", id).Error; err != nil {
		return nil, err
	}
	return merchant, nil
}

func CreateMerchant(merchant *models.Merchant) (interface{}, error) {
	if err := config.DB.Save(merchant).Error; err != nil {
		return nil, err
	}
	return merchant, nil
}

func UpdateMerchant(id string, merchant *models.Merchant) (interface{}, error) {
	var existingMerchant models.Merchant
	if err := config.DB.First(&existingMerchant, "merchant_id=?", id).Error; err != nil {
		return nil, err
	}
	existingMerchant.MerchantName = merchant.MerchantName
	existingMerchant.UserId = merchant.UserId
	if updateErr := config.DB.Save(&existingMerchant).Where("merchant_id=?", id).Error; updateErr != nil {
		return nil, updateErr
	}
	return existingMerchant, nil
}

func DeleteMerchant(id string) (interface{}, error) {
	var merchant models.Merchant
	if err := config.DB.First(&merchant, "merchant_id=?", id).Error; err != nil {
		return nil, err
	}
	if deleteErr := config.DB.Delete(&merchant).Where("merchant_id=?", id).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return merchant, nil
}
