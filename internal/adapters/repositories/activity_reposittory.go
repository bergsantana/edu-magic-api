package repositories

import (
	"github.com/bergsantana/edu-magic-api/internal/core/domain"
	"gorm.io/gorm"
)

type ActivityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) *ActivityRepository {
	return &ActivityRepository{db: db}
}

func (r *ActivityRepository) CreateActivity(activity *domain.Activity) error {
	return r.db.Create(activity).Error
}

func (r *ActivityRepository) GetActivitiesByUser(ownerID int64) (*[]domain.Activity, error) {
	var activities []domain.Activity
	if err := r.db.Where("owner_id = ?", ownerID).Find(&activities).Error; err != nil {
		return nil, err
	}
	return &activities, nil

}
