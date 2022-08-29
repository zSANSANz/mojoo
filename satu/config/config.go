package config

import (
	"Satu/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_Username"),
			os.Getenv("DB_Password"),
			os.Getenv("DB_Host"),
			os.Getenv("DB_Port"),
			os.Getenv("DB_Name"),
		)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func InitDBTest() {

	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			"root",
			"",
			"localhost",
			"3306",
			"mojoo",
		)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DropTable()
}

func InitialMigration() {
	DB.AutoMigrate(&models.Customer{}, &models.Product{}, &models.PaymentMethod{}, &models.Order{}, &models.OrderDetail{})
}

func DropTable() {
	DB.Migrator().DropTable(&models.Customer{}, &models.Product{}, &models.PaymentMethod{}, &models.Order{}, &models.OrderDetail{})
}
