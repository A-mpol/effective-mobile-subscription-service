package converter

import (
	"database/sql"
	"time"

	serviceModel "github.com/A-mpol/effective-mobile-subscription-service/internal/model"
	repoModel "github.com/A-mpol/effective-mobile-subscription-service/internal/repository/model"
)

func NewSubscriptionInfoToEntry(subscriptionInfo serviceModel.NewSubscriptionInfo) repoModel.NewSubscriptionEntry {
	return repoModel.NewSubscriptionEntry{
		UserId:      subscriptionInfo.UserId,
		ServiceName: subscriptionInfo.ServiceName,
		Price:       subscriptionInfo.Price,
		StartDate:   *subscriptionInfo.StartDate,
		EndDate:     endDateToEntry(subscriptionInfo.EndDate),
	}
}

func endDateToEntry(endDate *time.Time) sql.NullTime {
	if endDate == nil {
		return sql.NullTime{
			Valid: false,
		}
	}

	return sql.NullTime{
		Time:  *endDate,
		Valid: true,
	}
}

func ListSubscriptionEntriesToSubscription(subscriptionsEntries []repoModel.Subscription) []serviceModel.Subscription {
	subscriptions := []serviceModel.Subscription{}
	for _, subscriptionEntrie := range subscriptionsEntries {
		subscriptions = append(subscriptions, repoSubscriptionToService(subscriptionEntrie))
	}
	return subscriptions
}

func repoSubscriptionToService(subscription repoModel.Subscription) serviceModel.Subscription {
	return serviceModel.Subscription{
		SubscriptionId: subscription.SubscriptionId,
		UserId:         subscription.UserId,
		ServiceName:    subscription.ServiceName,
		Price:          subscription.Price,
		StartDate:      &subscription.StartDate,
		EndDate:        entryEndDateToModel(subscription.EndDate),
	}
}

func entryEndDateToModel(endDate sql.NullTime) *time.Time {
	if endDate.Valid {
		return &endDate.Time
	}

	return nil
}

func UpdateSubscriptionToEntry(updatedSubscription serviceModel.UpdatedSubscription) repoModel.UpdateSubscriptionEntry {
	return repoModel.UpdateSubscriptionEntry{
		SubscriptionId: updatedSubscription.SubscriptionId,
		UserId:         updatedSubscription.UserId,
		ServiceName:    updatedSubscription.ServiceName,
		Price:          updatedSubscription.Price,
		StartDate:      *updatedSubscription.StartDate,
		EndDate:        endDateToEntry(updatedSubscription.EndDate),
	}
}

func SubscriptionsFiltersToRepo(subscriptionsFilters serviceModel.SubscriptionsFilters) repoModel.SubscriptionsFilters {
	return repoModel.SubscriptionsFilters{
		StartDate:    *subscriptionsFilters.StartDate,
		EndDate:      endDateToEntry(subscriptionsFilters.EndDate),
		UserIds:      subscriptionsFilters.UserIds,
		ServiceNames: subscriptionsFilters.ServiceNames,
	}
}
