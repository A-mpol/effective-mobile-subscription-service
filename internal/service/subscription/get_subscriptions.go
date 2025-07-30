package subscription

import (
	"context"

	serviceModel "github.com/A-mpol/effective-mobile-subscription-service/internal/model"
	"github.com/google/uuid"
)

func (s *serviceImpl) GetSubscriptions(ctx context.Context, userID uuid.UUID) ([]serviceModel.Subscription, error) {
	list, err := s.repos.GetSubscriptions(ctx, userID)
	if err != nil {
		return nil, err
	}

	return list, nil
}
