package v1

import (
	"context"

	"github.com/A-mpol/effective-mobile-subscription-service/internal/converter"
	subscription_v1 "github.com/A-mpol/effective-mobile-subscription-service/pkg/proto/subscription/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *api) GetListSubscriptions(ctx context.Context, in *emptypb.Empty) (*subscription_v1.GetListSubscriptionsResponse, error) {
	list, err := a.service.GetListSubscriptions(ctx)
	if err != nil {
		return nil, err
	}

	return &subscription_v1.GetListSubscriptionsResponse{
		SubscriptionsList: converter.ListServiceSubscriptionToApiSubscription(list),
	}, nil
}
