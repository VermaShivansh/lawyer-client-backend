package models

import (
	"gorm.io/gorm"
)

type Court struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Type     string `gorm:"not null" json:"type"`
	Location string `gorm:"not null" json:"location"`
}

type Language struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
}

type PracticeArea struct {
	gorm.Model
	Name   string `gorm:"not null" json:"name"`
	AvgFee string `gorm:"not null" json:"avg_fee"`
}
