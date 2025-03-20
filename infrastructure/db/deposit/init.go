package deposit

import (
	"context"
	"fmt"
)

func (r *Repository) Init(ctx context.Context) error {
	tokensTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.tokens 
	(name varchar(50) UNIQUE, 
	address varchar(60) UNIQUE, 
	multiplicator float8,
	decimals int2,
	budget int8,
	meta bytea,
	active boolean,
	custom_score_table varchar(50) null,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP)`,
		r.db.Scheme(),
	)

	_, err := r.db.ExecContext(ctx, tokensTableQuery)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	claimsTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.claims 
	(
	id BIGSERIAL PRIMARY KEY,
	token varchar(50) not null, 
	user_id int8 not null, 
	amount int8 not null,
	tx_hash varchar(300),
	wallet varchar(100) not null,
	confirmed_at timestamp,
	created_at TIMESTAMP not null)`,
		r.db.Scheme(),
	)

	_, err = r.db.ExecContext(ctx, claimsTableQuery)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	replenishmentsTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.replenishments 
	(
	id BIGSERIAL PRIMARY KEY,
	token varchar(50) not null, 
	claim_id int8  null,
	user_id int8 not null, 
	wallet varchar(60) not null,
	amount int8 not null,
	comment varchar(60) not null,
	tx_hash varchar(100),
	tx_lt NUMERIC not null,
	created_at TIMESTAMP not null
	)`,
		r.db.Scheme(),
	)

	_, err = r.db.ExecContext(ctx, replenishmentsTableQuery)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	return nil
}
