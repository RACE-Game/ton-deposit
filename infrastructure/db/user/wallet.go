package user

import (
	"context"
	"fmt"
	"time"
)

func (r *Repository) SaveWallet(ctx context.Context, userID int64, wallet string) error {
	query := fmt.Sprintf(`INSERT INTO %s.wallets
	 (user_id, wallet,updated_at) VALUES ($1, $2,$3)`,
		r.db.Scheme(),
	)

	_, err := r.db.ExecContext(
		ctx,
		query,
		userID,
		wallet,
		time.Now().UTC(),
	)
	if err != nil {
		return fmt.Errorf("can't save wallet: %w", err)
	}

	return nil
}

func (r *Repository) UpdateWallet(ctx context.Context, userID int64, wallet string) error {
	query := fmt.Sprintf(`UPDATE %s.wallets SET wallet = $1, updated_at = $2 WHERE user_id = $3`,
		r.db.Scheme(),
	)

	_, err := r.db.ExecContext(
		ctx,
		query,
		wallet,
		time.Now().UTC(),
		userID,
	)
	if err != nil {
		return fmt.Errorf("can't update wallet: %w", err)
	}

	return nil
}

func (r *Repository) WalletExists(ctx context.Context, userID int64) (bool, error) {
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s.wallets WHERE user_id = $1)`,
		r.db.Scheme(),
	)

	var exists bool

	rows, err := r.db.QueryContext(
		ctx,
		query,
		userID,
	)
	if err != nil {
		return false, fmt.Errorf("can't check wallet existance: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&exists)
		if err != nil {
			return false, err
		}
	}

	return exists, nil
}

func (r *Repository) GetUserWallet(ctx context.Context, userID int64) (string, error) {
	query := fmt.Sprintf(`SELECT wallet FROM %s.wallets WHERE user_id = $1`,
		r.db.Scheme(),
	)

	var wallet string

	rows, err := r.db.QueryContext(
		ctx,
		query,
		userID,
	)
	if err != nil {
		return "", fmt.Errorf("can't get wallet: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&wallet)
		if err != nil {
			return "", err
		}
	}

	return wallet, nil
}
