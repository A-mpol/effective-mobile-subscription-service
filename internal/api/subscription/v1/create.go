package v1

import (
	"context"

	"github.com/A-mpol/effective-mobile-subscription-service/internal/converter"
	subscription_v1 "github.com/A-mpol/effective-mobile-subscription-service/pkg/proto/subscription/v1"
)

func (a *api) CreateSubscription(ctx context.Context, in *subscription_v1.CreateSubscriptionRequest) (*subscription_v1.CreateSubscriptionResponse, error) {
	newSubscriptionInfo := converter.CreateSubscriptionRequesToService(in)
	id, err := a.service.CreateSubscription(ctx, newSubscriptionInfo)
	if err != nil {
		return nil, err
	}

	return &subscription_v1.CreateSubscriptionResponse{
		SubscriptionId: id.String(),
	}, nil
}
