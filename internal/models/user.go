package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Email     string         `json:"email" gorm:"unique"`
	Name      string         `json:"name" gorm:"varchar(255)"`
	Role      string         `json:"role" gorm:"varchar(255)"`
	Password  string         `json:"-" gorm:"varchar(1024)"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

const ADMIN = "admin"
const MANAGER = "manager"
const USER = "user"
