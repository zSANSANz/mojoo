package db

import (
	"Satu/config"
	"Satu/models"
)

func GetOutlets() (interface{}, error) {
	var outlet []models.Outlet

	if err := config.DB.Find(&outlet).Error; err != nil {
		return nil, err
	}
	return outlet, nil
}

func GetOutletById(id string) (interface{}, error) {
	var outlet []models.Outlet

	if err := config.DB.First(&outlet, "outlet_id=?", id).Error; err != nil {
		return nil, err
	}
	return outlet, nil
}

func CreateOutlet(outlet *models.Outlet) (interface{}, error) {
	if err := config.DB.Save(outlet).Error; err != nil {
		return nil, err
	}
	return outlet, nil
}

func UpdateOutlet(id string, outlet *models.Outlet) (interface{}, error) {
	var existingOutlet models.Outlet
	if err := config.DB.First(&existingOutlet, "outlet_id=?", id).Error; err != nil {
		return nil, err
	}
	existingOutlet.OutletName = outlet.OutletName
	existingOutlet.MerchantId = outlet.MerchantId
	if updateErr := config.DB.Save(&existingOutlet).Where("outlet_id=?", id).Error; updateErr != nil {
		return nil, updateErr
	}
	return existingOutlet, nil
}

func DeleteOutlet(id string) (interface{}, error) {
	var outlet models.Outlet
	if err := config.DB.First(&outlet, "outlet_id=?", id).Error; err != nil {
		return nil, err
	}
	if deleteErr := config.DB.Delete(&outlet).Where("outlet_id=?", id).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return outlet, nil
}
