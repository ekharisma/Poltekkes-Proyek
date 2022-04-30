package entity

import (
	"time"
)

type Temperature struct {
	ID          uint `gorm:"primaryKey;autoincrement:true"`
	Timestamp   string
	Temperature string
	TimeCreated time.Time
}
