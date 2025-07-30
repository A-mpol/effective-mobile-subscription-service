package v1

import (
	"context"

	"github.com/A-mpol/effective-mobile-subscription-service/internal/converter"
	subscription_v1 "github.com/A-mpol/effective-mobile-subscription-service/pkg/proto/subscription/v1"
	"github.com/google/uuid"
)

func (a *api) GetSubscriptions(ctx context.Context, in *subscription_v1.GetSubscriptionsRequest) (*subscription_v1.GetSubscriptionsResponse, error) {
	list, err := a.service.GetSubscriptions(ctx, uuid.MustParse(in.GetUserId()))
	if err != nil {
		return nil, err
	}

	return &subscription_v1.GetSubscriptionsResponse{
		Subscriptions: converter.ListServiceSubscriptionToApiSubscription(list),
	}, nil
}
