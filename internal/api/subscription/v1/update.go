package v1

import (
	"context"

	"github.com/A-mpol/effective-mobile-subscription-service/internal/converter"
	subscription_v1 "github.com/A-mpol/effective-mobile-subscription-service/pkg/proto/subscription/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *api) UpdateSubscription(ctx context.Context, in *subscription_v1.UpdateSubscriptionRequest) (*emptypb.Empty, error) {
	updatedSubscription := converter.UpdateSubscriptionRequestToService(in)
	if err := a.service.UpdateSubscription(ctx, updatedSubscription); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
