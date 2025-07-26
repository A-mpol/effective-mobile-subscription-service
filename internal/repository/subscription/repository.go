package subscription

import (
	def "github.com/A-mpol/effective-mobile-subscription-service/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ def.SubscriptionRepository = (*repository)(nil)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repository {
	return &repository{
		db: db,
	}
}
