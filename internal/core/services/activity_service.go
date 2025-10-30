package services

import (
	"errors"

	"github.com/bergsantana/edu-magic-api/internal/adapters/repositories"
	"github.com/bergsantana/edu-magic-api/internal/core/domain"
)

type ActivityService struct {
	activityRepo repositories.ActivityRepository
}

func NewActivityService(activityRepo repositories.ActivityRepository) *ActivityService {
	return &ActivityService{activityRepo: activityRepo}
}

func (s *ActivityService) CreateActivity(activity *domain.Activity) error {
	if activity.OwnerId == 0 {
		return errors.New("owner ID must be set")
	}
	if len(activity.Title) == 0 {
		return errors.New("title cannot be empty")
	}
	if len(activity.WordSearch) == 0 {
		return errors.New("word search cannot be empty")
	}
	return s.activityRepo.CreateActivity(activity)
}

func (s *ActivityService) GetActivitiesByUser(ownerID int64) ([]*domain.Activity, error) {
	activities, err := s.activityRepo.GetActivitiesByUser(ownerID)
	if err != nil {
		return nil, err
	}
	var result []*domain.Activity
	for i := range *activities {
		result = append(result, &(*activities)[i])
	}
	return result, nil

}
