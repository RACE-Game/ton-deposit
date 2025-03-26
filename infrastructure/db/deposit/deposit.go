package deposit

import (
	"context"

	"github.com/google/uuid"
)

func (r *Repository) CreateDeposit(ctx context.Context, orderID uuid.UUID, userID int64, wallet, token string, amount uint64) error {
	query := `INSERT INTO deposits (order_id, user_id, wallet,token, amount) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.ExecContext(ctx, query, orderID, userID, token, amount)
	if err != nil {
		return err
	}

	return nil
}
