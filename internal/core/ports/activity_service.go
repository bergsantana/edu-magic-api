package ports

import "github.com/bergsantana/edu-magic-api/internal/core/domain"

type ActivityService interface {
	SaveActivity(title string, wordSearch [][]string) error
	GetActivitiesByUser(userID int64) ([]*domain.Activity, error)
}
