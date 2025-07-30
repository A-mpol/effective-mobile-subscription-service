package subscription

import (
	"context"
	"database/sql"

	serviceModel "github.com/A-mpol/effective-mobile-subscription-service/internal/model"
	"github.com/A-mpol/effective-mobile-subscription-service/internal/repository/converter"
	sq "github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *repository) UpdateSubscription(ctx context.Context, updatedSubscription serviceModel.UpdatedSubscription) error {
	updatedSubscriptionEntry := converter.UpdateSubscriptionToEntry(updatedSubscription)

	qb := sq.Update("subscriptions").
		PlaceholderFormat(sq.Dollar).
		Set("user_id", updatedSubscriptionEntry.UserId).
		Set("service_name", updatedSubscriptionEntry.ServiceName).
		Set("price", updatedSubscriptionEntry.Price).
		Set("start_date", updatedSubscriptionEntry.StartDate).
		Set("end_date", updatedSubscriptionEntry.EndDate).
		Where(sq.Eq{"id": updatedSubscriptionEntry.SubscriptionId}).
		Where(sq.Eq{"deleted_at": sql.NullTime{Valid: false}})

	query, args, err := qb.ToSql()
	if err != nil {
		return err
	}
	res, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return status.Error(codes.NotFound, "subscription not found or already deleted")
	}
	return nil
}
