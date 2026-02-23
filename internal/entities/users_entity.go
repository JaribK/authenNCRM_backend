package entities

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	User struct {
		Id        uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;not null"`
		Username  string         `gorm:"column:username;type:varchar(255);unique"`
		Password  string         `gorm:"column:password;type:varchar(255)"`
		Email     *string        `gorm:"column:email;type:varchar(255);unique"`
		FirstName string         `gorm:"column:first_name;type:varchar(255)"`
		LastName  string         `gorm:"column:last_name;type:varchar(255)"`
		RoleId    *uuid.UUID     `gorm:"column:role_id;type:uuid;index"`
		CreatedAt time.Time      `gorm:"column:created_at;"`
		UpdatedAt *time.Time     `gorm:"column:updated_at;"`
		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
		CreatedBy uuid.UUID      `gorm:"column:created_by;type:uuid"`
		UpdatedBy *uuid.UUID     `gorm:"column:updated_by;type:uuid"`
		DeletedBy *uuid.UUID     `gorm:"column:deleted_by;type:uuid"`

		Role      *Role       `gorm:"foreignKey:RoleId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
		PointLogs []PointLogs `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}

	UserRequest struct {
		Username        string  `json:"username"`
		Password        string  `json:"password"`
		ConfirmPassword string  `json:"confirmPassword"`
		FirstName       string  `json:"firstName"`
		LastName        string  `json:"lastName"`
		Email           *string `json:"email"`
		RoleId          *string `json:"roleId"`
		CreatedBy       string  `json:"-"`
		UpdatedBy       *string `json:"-"`
	}

	UserResponse struct {
		Id        uuid.UUID `json:"id"`
		Username  string    `json:"username"`
		Email     *string   `json:"email"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Role      *Role     `json:"role"`
		Points    int       `json:"points"`
	}
)

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	timeNew := time.Now()

	u.Id = uuid.New()
	u.CreatedAt = timeNew
	u.UpdatedAt = &timeNew

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}
