package repositories

import (
	"authenncrm/internal/entities"

	"gorm.io/gorm"
)

type (
	RoleRepository interface {
		GetRoleById(id string) (*entities.Role, error)
	}

	roleRepository struct {
		db *gorm.DB
	}
)

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) GetRoleById(id string) (*entities.Role, error) {
	var role entities.Role
	if err := r.db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}
