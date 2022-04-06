package model

import (
	"fmt"
	"log"

	"github.com/ekharisma/poltekkes-webservice/entity"
	"gorm.io/gorm"
)

type ITemperatureModel interface {
	StoreTemperature(temperature *entity.Temperature) error
	GetTemperatureByMonth(month, year int) ([]*entity.Temperature, error)
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

func (t TemperatureModel) GetTemperatureByMonth(month, year int) ([]*entity.Temperature, error) {
	var temperature []*entity.Temperature
	thisDate, toDate := constructDate(month, year)
	if err := t.db.Find(&temperature).Where("timestamp BETWEEN ? AND ?", thisDate, toDate).Error; err != nil {
		return nil, err
	}
	return temperature, nil
}

func constructDate(month, year int) (thisDate, toDate string) {
	if month == 12 {
		thisDate = fmt.Sprintf("%v-%v-01", year, month)
		toDate = fmt.Sprintf("%v-%v-01", year+1, month)
		return
	}
	thisDate = fmt.Sprintf("%v-%v-01", year, month)
	toDate = fmt.Sprintf("%v-%v-01", year, month+1)
	return
}
