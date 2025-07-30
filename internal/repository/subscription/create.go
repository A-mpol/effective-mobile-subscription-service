package subscription

import (
	"context"

	serviceModel "github.com/A-mpol/effective-mobile-subscription-service/internal/model"
	"github.com/A-mpol/effective-mobile-subscription-service/internal/repository/converter"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (r *repository) CreateSubscription(ctx context.Context, subscriptionInfo serviceModel.NewSubscriptionInfo) (uuid.UUID, error) {
	subscriptionEntry := converter.NewSubscriptionInfoToEntry(subscriptionInfo)

	subscriptionUUID := uuid.New()

	qb := sq.Insert("subscriptions").
		PlaceholderFormat(sq.Dollar).
		Columns("id", "user_id", "service_name", "price", "start_date", "end_date").
		Values(
			subscriptionUUID,
			subscriptionEntry.UserId,
			subscriptionEntry.ServiceName,
			subscriptionEntry.Price,
			subscriptionEntry.StartDate,
			subscriptionEntry.EndDate,
		)

	query, args, err := qb.ToSql()
	if err != nil {
		return uuid.UUID{}, err
	}

	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		return uuid.UUID{}, err
	}

	return subscriptionUUID, nil
}
