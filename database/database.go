package database

import (
	"app/model"
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func Connect() {
	var err error
	database, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&model.Conversion{})
}

func Save(conversion model.Conversion) {
	database.Create(&conversion)
}

func GetConversions() ([]model.Conversion, error) {
	var conversions []model.Conversion
	result := database.Find(&conversions)
	return conversions, result.Error
}

func GetConversion(id string) (model.Conversion, error) {
	var conversion model.Conversion

	if err := database.First(&conversion, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return model.Conversion{}, errors.New("No record found for specified ID")
	}

	return conversion, nil
}
