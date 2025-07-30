package converter

import (
	"time"

	serviceModel "github.com/A-mpol/effective-mobile-subscription-service/internal/model"
	subscription_v1 "github.com/A-mpol/effective-mobile-subscription-service/pkg/proto/subscription/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateSubscriptionRequesToService(in *subscription_v1.CreateSubscriptionRequest) serviceModel.NewSubscriptionInfo {
	return serviceModel.NewSubscriptionInfo{
		UserId:      in.UserId,
		ServiceName: in.ServiceName,
		Price:       in.Price,
		StartDate:   convertProtoTimestamp(in.StartDate),
		EndDate:     convertProtoTimestamp(in.EndDate),
	}
}

func convertProtoTimestamp(ts *timestamppb.Timestamp) *time.Time {
	if ts == nil {
		return nil
	}
	t := ts.AsTime()
	return &t
}

func ListServiceSubscriptionToApiSubscription(serviceSubscriptions []serviceModel.Subscription) []*subscription_v1.Subscription {
	subscriptions := []*subscription_v1.Subscription{}
	for _, serviceSubscription := range serviceSubscriptions {
		subscriptions = append(subscriptions, serviceSubscriptionToApi(serviceSubscription))
	}
	return subscriptions
}

func serviceSubscriptionToApi(subscription serviceModel.Subscription) *subscription_v1.Subscription {
	return &subscription_v1.Subscription{
		SubscriptionId: subscription.SubscriptionId,
		UserId:         subscription.UserId,
		ServiceName:    subscription.ServiceName,
		Price:          subscription.Price,
		StartDate:      convertTimeToProto(subscription.StartDate),
		EndDate:        convertTimeToProto(subscription.EndDate),
	}
}

func convertTimeToProto(ts *time.Time) *timestamppb.Timestamp {
	if ts == nil {
		return nil
	}
	return timestamppb.New(*ts)
}

func GetTotalCostSubscriptionsRequestToService(in *subscription_v1.GetTotalCostSubscriptionsRequest) serviceModel.SubscriptionsFilters {
	return serviceModel.SubscriptionsFilters{
		StartDate:    convertProtoTimestamp(in.StartDate),
		EndDate:      convertProtoTimestamp(in.EndDate),
		UserIds:      in.UserIds,
		ServiceNames: in.ServiceNames,
	}
}

func UpdateSubscriptionRequestToService(in *subscription_v1.UpdateSubscriptionRequest) serviceModel.UpdatedSubscription {
	return serviceModel.UpdatedSubscription{
		SubscriptionId: in.SubscriptionId,
		UserId:         in.UserId,
		ServiceName:    in.ServiceName,
		Price:          in.Price,
		StartDate:      convertProtoTimestamp(in.StartDate),
		EndDate:        convertProtoTimestamp(in.EndDate),
	}
}
