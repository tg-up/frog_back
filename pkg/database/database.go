package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"icecreambash/tgup_backend/internal/config"
	"icecreambash/tgup_backend/internal/models"
	"time"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow", config.GlobalConfig.DBHost, config.GlobalConfig.DBUser, config.GlobalConfig.DBPassword, config.GlobalConfig.DBName, config.GlobalConfig.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Europe/Moscow")
			return time.Now().In(ti)
		},
	})
	if err != nil {
		return err
	}
	DB = db

	DB.AutoMigrate(&models.Platform{})
	DB.AutoMigrate(&models.PlatformServices{})
	DB.AutoMigrate(&models.ServiceField{})
	DB.AutoMigrate(&models.User{})
	return nil
}
