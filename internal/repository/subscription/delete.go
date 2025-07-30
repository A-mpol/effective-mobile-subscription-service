package subscription

import (
	"context"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (r *repository) DeleteSubscription(ctx context.Context, subscriptionID uuid.UUID) error {
	qb := sq.Update("subscriptions").
		PlaceholderFormat(sq.Dollar).
		Set("deleted_at", time.Now()).
		Where(sq.And{
			sq.Eq{"id": subscriptionID},
			sq.Eq{"deleted_at": nil},
		})

	query, args, err := qb.ToSql()
	if err != nil {
		return err
	}

	log.Println(query)
	log.Println(args)
	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		return err
	}

	return nil
}
