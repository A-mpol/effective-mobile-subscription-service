package subscription

import (
	"context"

	serviceModel "github.com/A-mpol/effective-mobile-subscription-service/internal/model"
	"github.com/google/uuid"
)

func (s *serviceImpl) CreateSubscription(ctx context.Context, subscriptionInfo serviceModel.NewSubscriptionInfo) (uuid.UUID, error) {
	id, err := s.repos.CreateSubscription(ctx, subscriptionInfo)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}
