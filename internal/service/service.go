package service

import (
	"context"

	serviceModel "github.com/A-mpol/effective-mobile-subscription-service/internal/model"
	"github.com/google/uuid"
)

type SubscriptionService interface {
	CreateSubscription(ctx context.Context, subscriptionInfo serviceModel.NewSubscriptionInfo) (uuid.UUID, error)
	GetSubscriptions(ctx context.Context, userID uuid.UUID) ([]serviceModel.Subscription, error)
	UpdateSubscription(ctx context.Context, updatedSubscription serviceModel.UpdatedSubscription) error
	DeleteSubscription(ctx context.Context, subscriptionID uuid.UUID) error
	GetTotalCostSubscriptions(ctx context.Context, filters serviceModel.SubscriptionsFilters) (totalCost int64, err error)
	GetListSubscriptions(ctx context.Context) ([]serviceModel.Subscription, error)
}
