package user

import (
	"context"
	"fmt"
)

func (r *Repository) Init(ctx context.Context) error {
	usersTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.users (
	telegram_id bigint PRIMARY KEY, 
	tg_user_name  varchar(100),
	first_name varchar(100),
	last_name varchar(100),
	language_code varchar(100),
	is_premium boolean,
	chat_id bigint,  
	start_message_id bigint,
	last_message_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
)`, r.db.Scheme())

	_, err := r.db.ExecContext(ctx, usersTableQuery)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	dataTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.user_data (
		user_id bigint PRIMARY KEY, 
		data bytea,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
	)`, r.db.Scheme())

	_, err = r.db.ExecContext(ctx, dataTableQuery)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	walletTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.wallets (
		user_id bigint PRIMARY KEY, 
		wallet varchar(100),
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
	)`, r.db.Scheme())

	_, err = r.db.ExecContext(ctx, walletTableQuery)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	notificationTableQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.notification_results (
		user_id bigint PRIMARY KEY,
		result varchar(100),
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`, r.db.Scheme())

	_, err = r.db.ExecContext(ctx, notificationTableQuery)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	return nil
}
