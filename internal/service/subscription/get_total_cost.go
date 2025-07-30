package subscription

import (
	"context"

	serviceModel "github.com/A-mpol/effective-mobile-subscription-service/internal/model"
)

func (s *serviceImpl) GetTotalCostSubscriptions(ctx context.Context, filters serviceModel.SubscriptionsFilters) (totalCost int64, err error) {
	totalCost, err = s.repos.GetTotalCostSubscriptions(ctx, filters)
	if err != nil {
		return 0, err
	}

	return totalCost, nil
}
