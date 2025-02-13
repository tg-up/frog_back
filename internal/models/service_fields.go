package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type JSONB []interface{}

// Value Marshal
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		fmt.Println(ok)
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

type ServiceField struct {
	gorm.Model
	ServiceID    uint     `json:"service_id"`
	FieldName    string   `json:"field_name"`
	FieldSlug    string   `json:"field_slug"`
	FieldType    string   `json:"field_type"` //text,number,radio,checkbox
	DefaultValue string   `json:"default_value"`
	MinValue     uint8    `json:"min_value"`
	MaxValue     uint8    `json:"max_value"`
	CostType     string   `json:"cost_type"` //coefficient|base
	Options      JSONB    `json:"options" gorm:"type:jsonb;serializer:json"`
	TempOptions  []Option `json:"-" gorm:"-"`
	Size         string   `json:"size"` //fit,grow
}

type Option struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}
