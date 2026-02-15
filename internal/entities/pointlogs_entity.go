package entities

import "github.com/google/uuid"

type (
	PointLogs struct {
		Id       uuid.UUID `gorm:"column:id;type:uuid;primaryKey;not null"`
		UserId   uuid.UUID `gorm:"column:user_id;type:uuid;index"`
		Points   int       `gorm:"column:points;default:0"`
		Activity string    `gorm:"column:activity;type:varchar(255)"`

		User User `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
