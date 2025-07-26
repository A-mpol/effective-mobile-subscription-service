package repository

import (
	"context"

	serviceModel "github.com/A-mpol/effective-mobile-subscription-service/internal/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SubscriptionRepository interface {
	CreateSubscription(ctx context.Context, subscriptionInfo serviceModel.NewSubscriptionInfo) (subscriptionId string, err error)
	GetSubscriptions(ctx context.Context, userID string) ([]serviceModel.Subscription, error)
	UpdateSubscription(ctx context.Context, updatedSubscription serviceModel.UpdatedSubscription) (*emptypb.Empty, error)
	DeleteSubscription(ctx context.Context, subscriptionID string) (*emptypb.Empty, error)
	GetTotalCostSubscriptions(ctx context.Context, filters serviceModel.SubscriptionsFilters) (totalCost int64, err error)
	GetListSubscriptions(ctx context.Context, in *emptypb.Empty) ([]serviceModel.Subscription, error)
}
