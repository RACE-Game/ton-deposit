package deposit

import (
	"context"
	"fmt"
)

func (r *Repository) Init(ctx context.Context) error {
	ordersTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.orders 
	id UUID PRIMARY KEY,
	token varchar(50) not null,
	user_id int8 not null,
	amount int8 not null,
	wallet varchar(60) not null,
	tx_hash varchar(100),
	created_at TIMESTAMP not null default now()`,
		r.db.Scheme(),
	)

	_, err := r.db.ExecContext(ctx, ordersTableQuery)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	depositsTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.deposits 
	(
	id BIGSERIAL PRIMARY KEY,
	order_id UUID not null,
	user_id int8 not null,
	wallet varchar(60) not null,
	token varchar(50) not null,
	amount int8 not null, 
	created_at TIMESTAMP not null default now()
	)`,
		r.db.Scheme(),
	)

	_, err = r.db.ExecContext(ctx, depositsTableQuery)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	return nil
}
