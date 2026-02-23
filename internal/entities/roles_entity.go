package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Role struct {
		Id        uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;not null"`
		Name      string         `gorm:"column:name;type:varchar(255);unique"`
		Level     int            `gorm:"column:level;type:int;default:0"`
		CreatedAt time.Time      `gorm:"column:created_at;"`
		UpdatedAt *time.Time     `gorm:"column:updated_at;"`
		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
		CreatedBy uuid.UUID      `gorm:"column:created_by;type:uuid"`
		UpdatedBy *uuid.UUID     `gorm:"column:updated_by;type:uuid"`
		DeletedBy *uuid.UUID     `gorm:"column:deleted_by;type:uuid"`

		Users []User `gorm:"foreignKey:RoleId"`
	}

	RoleRequest struct {
		Name  string `json:"name"`
		Level int    `json:"level"`
	}
)
