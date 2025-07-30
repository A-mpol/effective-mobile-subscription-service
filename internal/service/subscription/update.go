package subscription

import (
	"context"

	serviceModel "github.com/A-mpol/effective-mobile-subscription-service/internal/model"
)

func (s *serviceImpl) UpdateSubscription(ctx context.Context, updatedSubscription serviceModel.UpdatedSubscription) error {
	if err := s.repos.UpdateSubscription(ctx, updatedSubscription); err != nil {
		return err
	}

	return nil
}
