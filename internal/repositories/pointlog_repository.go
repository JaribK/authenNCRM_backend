package repositories

import (
	"authenncrm/internal/entities"

	"gorm.io/gorm"
)

type (
	PointLogsRepository interface {
		GetPointLogsByUserId(userId string) ([]entities.PointLogs, error)
		CreatePointLog(log *entities.PointLogs) error
	}

	pointLogsRepository struct {
		db *gorm.DB
	}
)

func NewPointLogsRepository(db *gorm.DB) PointLogsRepository {
	return &pointLogsRepository{db: db}
}

func (r *pointLogsRepository) GetPointLogsByUserId(userId string) ([]entities.PointLogs, error) {
	var logs []entities.PointLogs
	if err := r.db.Where("user_id = ?", userId).Find(&logs).Error; err != nil {
		return nil, err
	}

	return logs, nil
}

func (r *pointLogsRepository) CreatePointLog(log *entities.PointLogs) error {
	if err := r.db.Create(log).Error; err != nil {
		return err
	}

	return nil
}
