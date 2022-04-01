package entity

import (
	"gorm.io/gorm"
)

type Temperature struct {
	gorm.Model
	Time        int64
	Temperature float64
}
