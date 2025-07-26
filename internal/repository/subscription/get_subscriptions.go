package subscription

import (
	"context"
	"database/sql"

	serviceModel "github.com/A-mpol/effective-mobile-subscription-service/internal/model"
	"github.com/A-mpol/effective-mobile-subscription-service/internal/repository/converter"
	repoModel "github.com/A-mpol/effective-mobile-subscription-service/internal/repository/model"
	sq "github.com/Masterminds/squirrel"
)

func (r *repository) GetSubscriptions(ctx context.Context, userID string) ([]serviceModel.Subscription, error) {
	qb := sq.Select("id", "service_name", "price", "start_date", "end_date").
		From("subscriptions").
		Where(sq.Eq{"user_id": userID}).
		Where(sq.Eq{"deleted_at": sql.NullTime{Valid: false}})

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	subscriptionsEntries := []repoModel.Subscription{}
	for rows.Next() {
		subscription := repoModel.Subscription{}
		err = rows.Scan(
			&subscription.SubscriptionId,
			&subscription.ServiceName,
			&subscription.Price,
			&subscription.StartDate,
			&subscription.EndDate,
		)
		if err != nil {
			return nil, err
		}
		subscriptionsEntries = append(subscriptionsEntries, subscription)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return converter.ListSubscriptionEntriesToSubscription(subscriptionsEntries), nil
}
