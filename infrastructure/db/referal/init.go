package referal

import (
	"context"
	"fmt"
)

func (r *Repository) Init(ctx context.Context) error {
	referalsTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.referals (
	referrer_user_id BIGINT, referal_user_id BIGINT, 
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	UNIQUE(referrer_user_id, referal_user_id))`, r.db.Scheme())

	_, err := r.db.ExecContext(ctx, referalsTableQuery)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	return nil
}
