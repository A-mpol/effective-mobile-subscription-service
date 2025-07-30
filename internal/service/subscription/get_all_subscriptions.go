package subscription

import (
	"context"

	serviceModel "github.com/A-mpol/effective-mobile-subscription-service/internal/model"
)

func (s *serviceImpl) GetListSubscriptions(ctx context.Context) ([]serviceModel.Subscription, error) {
	list, err := s.repos.GetListSubscriptions(ctx)
	if err != nil {
		return nil, err
	}

	return list, nil
}
