package models

import "gorm.io/gorm"

type Platform struct {
	gorm.Model
	Name     string `json:"name"`
	Slug     string `json:"slug" gorm:"unique"`
	Services []PlatformServices
}

type PlatformServices struct {
	gorm.Model
	Name       string
	Slug       string
	PlatformID uint `json:"platform_id"`
}
