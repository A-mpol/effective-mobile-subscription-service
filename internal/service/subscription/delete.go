package subscription

import (
	"context"

	"github.com/google/uuid"
)

func (s *serviceImpl) DeleteSubscription(ctx context.Context, subscriptionID uuid.UUID) error {
	if err := s.repos.DeleteSubscription(ctx, subscriptionID); err != nil {
		return err
	}

	return nil
}
