package model

import (
	"github.com/ekharisma/poltekkes-webservice/entity"
	"gorm.io/gorm"
	"log"
)

type ITemperatureModel interface {
	StoreTemperature(temperature *entity.Temperature) error
	GetTemperatureByMonth() ([]*entity.Temperature, error)
}

type TemperatureModel struct {
	db *gorm.DB
}

func CreateTemperatureModel(db *gorm.DB) ITemperatureModel {
	o := TemperatureModel{db: db}
	if err := o.db.AutoMigrate(&entity.Temperature{}); err != nil {
		log.Panicln("Error migrating temperature entity : ", err.Error())
	}
	return &o
}

func (t TemperatureModel) StoreTemperature(temperature *entity.Temperature) error {
	if err := t.db.Create(temperature).Error; err != nil {
		return err
	}
	return nil
}

func (t TemperatureModel) GetTemperatureByMonth() ([]*entity.Temperature, error) {
	var temperature []*entity.Temperature
	if err := t.db.Find(&temperature).Error; err != nil {
		return nil, err
	}
	return temperature, nil
}
