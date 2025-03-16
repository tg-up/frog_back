package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const (
	START       string = "start"
	IN_PROGRESS string = "in_progress"
	SUCCESS     string = "success"
	FAILURE     string = "failure"
)

type StatusOrder string

type Order struct {
	ID        uuid.UUID        `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	ServiceID int              `json:"service_id"`
	UserID    uuid.UUID        `json:"user_id"`
	Link      string           `json:"link"`
	Count     uint             `json:"count"`
	Reserved  bool             `json:"reserved" gorm:"default:false"`
	Status    StatusOrder      `json:"status" sql:"type:ENUM('start', 'in_progress', 'success','FAILURE')" gorm:"column:status"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"deleted_at"`
	Service   PlatformServices `gorm:"foreignKey:ServiceID"`
	User      User             `gorm:"foreignKey:UserID"`
}
