package model

import (
	"time"
)

type NewSubscriptionInfo struct {
	UserId      string
	ServiceName string
	Price       int64
	StartDate   *time.Time
	EndDate     *time.Time
}

type Subscription struct {
	SubscriptionId string
	UserId         string
	ServiceName    string
	Price          int64
	StartDate      *time.Time
	EndDate        *time.Time
}

type SubscriptionStatus int32

const (
	SubscriptionStatus_SUBSCRIPTION_STATUS_UNSPECIFIED SubscriptionStatus = 0
	SubscriptionStatus_SUBSCRIPTION_STATUS_ACTIVE      SubscriptionStatus = 1
	SubscriptionStatus_SUBSCRIPTION_STATUS_PAUSED      SubscriptionStatus = 2
	SubscriptionStatus_SUBSCRIPTION_STATUS_CANCELLED   SubscriptionStatus = 3
	SubscriptionStatus_SUBSCRIPTION_STATUS_EXPIRED     SubscriptionStatus = 4
)

type UpdatedSubscription struct {
	SubscriptionId string
	UserId         string
	ServiceName    string
	Price          int64
	StartDate      *time.Time
	EndDate        *time.Time
}

type SubscriptionsFilters struct {
	StartDate    *time.Time
	EndDate      *time.Time
	UserIds      []string
	ServiceNames []string
}
