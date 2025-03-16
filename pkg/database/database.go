package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"icecreambash/tgup_backend/internal/configs"
	"icecreambash/tgup_backend/internal/models"
	"time"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", configs.GlobalConfig.DBHost, configs.GlobalConfig.DBUser, configs.GlobalConfig.DBPassword, configs.GlobalConfig.DBName, configs.GlobalConfig.DBPort)
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
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Order{})
	return nil
}
