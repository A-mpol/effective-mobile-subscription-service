package subscription

import (
	"context"
	"database/sql"

	serviceModel "github.com/A-mpol/effective-mobile-subscription-service/internal/model"
	"github.com/A-mpol/effective-mobile-subscription-service/internal/repository/converter"
	sq "github.com/Masterminds/squirrel"
)

func (r *repository) GetTotalCostSubscriptions(ctx context.Context, filters serviceModel.SubscriptionsFilters) (totalCost int64, err error) {
	filtersEntry := converter.SubscriptionsFiltersToRepo(filters)

	qb := sq.Select("COALESCE(SUM(price), 0)").
		PlaceholderFormat(sq.Dollar).
		From("subscriptions").
		Where(sq.Eq{"deleted_at": sql.NullTime{Valid: false}})

	if len(filtersEntry.UserIds) > 0 {
		qb = qb.Where(sq.Eq{"user_id": filtersEntry.UserIds})
	}
	if len(filtersEntry.ServiceNames) > 0 {
		qb = qb.Where(sq.Eq{"service_name": filtersEntry.ServiceNames})
	}
	qb = qb.Where(sq.GtOrEq{"start_date": filtersEntry.StartDate})
	if filtersEntry.EndDate.Valid {
		qb = qb.Where(sq.LtOrEq{"end_date": filtersEntry.EndDate})
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return 0, err
	}

	row := r.db.QueryRow(ctx, query, args...)
	if err := row.Scan(&totalCost); err != nil {
		return 0, err
	}
	return totalCost, nil
}
