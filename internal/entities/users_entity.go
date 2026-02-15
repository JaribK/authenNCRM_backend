package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	User struct {
		Id        uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;not null"`
		Username  string         `gorm:"column:username;type:varchar(255);unique"`
		Password  string         `gorm:"column:password;type:varchar(255)"`
		Email     string         `gorm:"column:email;type:varchar(255);unique"`
		RoleId    uuid.UUID      `gorm:"column:role_id;type:uuid;index"`
		CreatedAt time.Time      `gorm:"column:created_at;"`
		UpdatedAt *time.Time     `gorm:"column:updated_at;"`
		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
		CreatedBy uuid.UUID      `gorm:"column:created_by;type:uuid"`
		UpdatedBy *uuid.UUID     `gorm:"column:updated_by;type:uuid"`
		DeletedBy *uuid.UUID     `gorm:"column:deleted_by;type:uuid"`

		Role      Role        `gorm:"foreignKey:RoleId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
		PointLogs []PointLogs `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
