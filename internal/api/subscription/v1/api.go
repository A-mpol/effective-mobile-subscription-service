package v1

import (
	"github.com/A-mpol/effective-mobile-subscription-service/internal/service"
	subscription_v1 "github.com/A-mpol/effective-mobile-subscription-service/pkg/proto/subscription/v1"
)

type api struct {
	subscription_v1.UnimplementedSubscriptionServiceServer

	service service.SubscriptionService
}

func NewApi(service service.SubscriptionService) *api {
	return &api{
		service: service,
	}
}
