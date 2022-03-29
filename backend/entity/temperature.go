package entity

import (
	"gorm.io/gorm"
	"time"
)

type Temperature struct {
	gorm.Model
	time        time.Time
	temperature float64
}
