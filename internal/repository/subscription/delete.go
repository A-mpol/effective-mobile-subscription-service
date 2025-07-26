package subscription

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (r *repository) DeleteSubscription(ctx context.Context, subscriptionID string) (*emptypb.Empty, error) {
	qb := sq.Update("subscriptions").
		Set("deleted_at", sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}).
		Where(sq.Eq{"id": subscriptionID}).
		Where(sq.Eq{"deleted_at": sql.NullTime{Valid: false}})

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
