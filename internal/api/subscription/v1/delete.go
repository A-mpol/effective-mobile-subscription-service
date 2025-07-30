package v1

import (
	"context"

	subscription_v1 "github.com/A-mpol/effective-mobile-subscription-service/pkg/proto/subscription/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *api) DeleteSubscription(ctx context.Context, in *subscription_v1.DeleteSubscriptionRequest) (*emptypb.Empty, error) {
	if err := a.service.DeleteSubscription(ctx, uuid.MustParse(in.GetSubscriptionId())); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
