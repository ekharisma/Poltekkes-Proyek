package entity

import (
	"time"
)

type Temperature struct {
	//gorm.Model
	ID          uint `gorm:"primaryKey;autoincrement:true"`
	Timestamp   string
	Temperature float64
	TimeCreated time.Time
}
