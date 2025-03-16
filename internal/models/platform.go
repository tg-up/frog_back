package models

import "gorm.io/gorm"

type Platform struct {
	gorm.Model
	Name     string             `json:"name"`
	Slug     string             `json:"slug" gorm:"unique"`
	Services []PlatformServices `gorm:"foreignKey:PlatformID"`
}

type PlatformServices struct {
	gorm.Model
	Name          string  `json:"name"`
	Slug          string  `json:"slug"`
	MinCount      uint    `json:"min_count"`
	MaxCount      uint    `json:"max_count"`
	Drip          bool    `json:"drip"`
	Amount        float64 `json:"amount"`
	AmountAbility uint    `json:"amount_ability"`
	PlatformID    uint    `json:"platform_id"`
}
