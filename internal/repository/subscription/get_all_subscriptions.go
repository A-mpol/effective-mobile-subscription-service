package subscription

import (
	"context"
	"database/sql"

	serviceModel "github.com/A-mpol/effective-mobile-subscription-service/internal/model"
	"github.com/A-mpol/effective-mobile-subscription-service/internal/repository/converter"
	repoModel "github.com/A-mpol/effective-mobile-subscription-service/internal/repository/model"
	sq "github.com/Masterminds/squirrel"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (r *repository) GetListSubscriptions(ctx context.Context, in *emptypb.Empty) ([]serviceModel.Subscription, error) {
	qb := sq.Select("*").
		From("subscriptions").
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
			&subscription.UserId,
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
