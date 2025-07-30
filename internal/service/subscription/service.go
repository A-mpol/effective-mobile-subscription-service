package subscription

import (
	"github.com/A-mpol/effective-mobile-subscription-service/internal/repository"
	"github.com/A-mpol/effective-mobile-subscription-service/internal/service"
)

var _ service.SubscriptionService = (*serviceImpl)(nil)

type serviceImpl struct {
	repos repository.SubscriptionRepository
}

func NewService(repos repository.SubscriptionRepository) *serviceImpl {
	return &serviceImpl{
		repos: repos,
	}
}
