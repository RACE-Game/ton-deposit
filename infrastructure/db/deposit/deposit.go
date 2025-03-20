package deposit

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (r *Repository) Order(ctx context.Context, token string, userID int64, amount int64, wallet string) (id uuid.UUID, err error) {
	query := fmt.Sprintf(`INSERT INTO %s.orders 
	(token, user_id, amount, wallet, created_at) VALUES ($1, $2, $3, $4, $5)
	returning id`,
		r.db.Scheme(),
	)

	row, err := r.db.QueryContext(ctx, query, token, userID, amount, wallet, time.Now())
	if err != nil {
		return uuid.Nil, fmt.Errorf("can't save claim: %w", err)
	}

	defer row.Close()

	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("can't scan created claim id: %w", err)
	}

	return id, nil
}

func (r *Repository) GetByUserID(ctx context.Context, userID int64) (claims []model.ClaimWithdrowal, err error) {
	query := fmt.Sprintf(`SELECT id, 
	token,user_id,amount, tx_hash,confirmed_at, wallet,created_at 
	FROM %s.claims WHERE user_id = $1`,
		r.db.Scheme(),
	)

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("can't get claims by user id: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var claim Claim
		err = rows.Scan(
			&claim.ID,
			&claim.Token,
			&claim.UserID,
			&claim.Amount,
			&claim.TxHash,
			&claim.ConfirmedAt,
			&claim.Wallet,
			&claim.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("can't scan claim: %w", err)
		}

		c := model.ClaimWithdrowal{
			ID:     claim.ID,
			Token:  claim.Token,
			UserID: claim.UserID,
			Amount: claim.Amount,
			Wallet: claim.Wallet,
			TX:     claim.TxHash,
		}

		if claim.ConfirmedAt.Valid {
			c.ConfirmedAt = &claim.ConfirmedAt.Time
		}

		if claim.CreatedAt.Valid {
			c.CreatedAt = claim.CreatedAt.Time
		}

		claims = append(claims, c)
	}

	return claims, nil
}

func (r *Repository) GetByID(ctx context.Context, userID int64) (claim model.ClaimWithdrowal, err error) {
	query := fmt.Sprintf(`SELECT id, 
	token,user_id,amount, tx_hash,confirmed_at,wallet, created_at 
	FROM %s.claims WHERE id = $1`,
		r.db.Scheme(),
	)

	row := r.db.QueryRowContext(ctx, query, userID)
	err = row.Scan(
		&claim.ID,
		&claim.Token,
		&claim.UserID,
		&claim.Amount,
		&claim.TX,
		&claim.ConfirmedAt,
		&claim.Wallet,
		&claim.CreatedAt,
	)
	if err != nil {
		return claim, fmt.Errorf("can't get claim by id: %w", err)
	}

	return claim, nil
}

func (r *Repository) SetTXHash(ctx context.Context, claimID int64, txHash string) error {
	query := fmt.Sprintf(`UPDATE %s.claims SET tx_hash = $1 WHERE id = $2`,
		r.db.Scheme(),
	)

	_, err := r.db.ExecContext(ctx, query, txHash, claimID)
	if err != nil {
		return fmt.Errorf("can't set tx hash: %w", err)
	}

	return nil
}

func (r *Repository) Confirm(ctx context.Context, claimID int64, txHash string) error {
	query := fmt.Sprintf(`UPDATE %s.claims SET confirmed_at = $1 WHERE tx_hash = $2`,
		r.db.Scheme(),
	)

	_, err := r.db.ExecContext(ctx, query, time.Now(), txHash)
	if err != nil {
		return fmt.Errorf("can't confirm claim: %w", err)
	}

	return nil
}

func (r *Repository) SetTx(ctx context.Context, txHash string, id int64) error {
	query := fmt.Sprintf(`UPDATE %s.claims SET tx_hash = $1 WHERE id = $2`,
		r.db.Scheme(),
	)

	_, err := r.db.ExecContext(ctx, query, txHash, id)
	if err != nil {
		return fmt.Errorf("can't set tx hash: %w", err)
	}

	return nil
}

func (r *Repository) GetByTx(ctx context.Context, txHash string) (claim model.ClaimWithdrowal, err error) {
	query := fmt.Sprintf(`SELECT id, 
	token,user_id,amount, tx_hash,confirmed_at,wallet, created_at
	 FROM %s.claims WHERE tx_hash = $1`,
		r.db.Scheme(),
	)

	row := r.db.QueryRowContext(ctx, query, txHash)
	err = row.Scan(
		&claim.ID,
		&claim.Token,
		&claim.UserID,
		&claim.Amount,
		&claim.TX,
		&claim.ConfirmedAt,
		&claim.Wallet,
		&claim.CreatedAt,
	)
	if err != nil {
		return claim, fmt.Errorf("can't get claim by tx: %w", err)
	}

	return claim, nil
}
