package v1

import (
	"context"

	"github.com/A-mpol/effective-mobile-subscription-service/internal/converter"
	subscription_v1 "github.com/A-mpol/effective-mobile-subscription-service/pkg/proto/subscription/v1"
)

func (a *api) GetTotalCostSubscriptions(ctx context.Context, in *subscription_v1.GetTotalCostSubscriptionsRequest) (*subscription_v1.GetTotalCostSubscriptionsResponse, error) {
	subscriptionsFilters := converter.GetTotalCostSubscriptionsRequestToService(in)

	totalCost, err := a.service.GetTotalCostSubscriptions(ctx, subscriptionsFilters)
	if err != nil {
		return nil, err
	}

	return &subscription_v1.GetTotalCostSubscriptionsResponse{
		TotalCost: totalCost,
	}, nil
}
