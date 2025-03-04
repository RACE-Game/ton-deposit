package referal

import (
	"context"
	"fmt"
	"time"

	"github.com/RACE-Game/ton-deposit-service/infrastructure/db"
	"github.com/RACE-Game/ton-deposit-service/internal/domain/telegram"
)

type Repository struct {
	db db.Database
}

func New(db db.Database) (*Repository, error) {

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}

	return &Repository{db: db}, nil
}

func (r *Repository) Check(ctx context.Context, referrerUserID, referalUserID int64) (bool, error) {
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1
	FROM %s.referals WHERE referrer_user_id = $1 AND referal_user_id = $2)`, r.db.Scheme())
	rows, err := r.db.QueryContext(
		ctx,
		query,
		referrerUserID, referalUserID,
	)
	if err != nil {
		return false, fmt.Errorf("can't check referals existance: %w", err)
	}

	defer rows.Close()

	var exist bool

	for rows.Next() {
		err := rows.Scan(&exist)
		if err != nil {
			return false, err
		}
	}

	return exist, nil
}

func (r *Repository) Save(ctx context.Context, referrerUserID int64, referalUserID int64) error {
	query := fmt.Sprintf(`INSERT INTO %s.referals (referrer_user_id, referal_user_id) 
	VALUES ($1, $2) ON CONFLICT(referrer_user_id, referal_user_id) DO NOTHING`,
		r.db.Scheme())

	_, err := r.db.ExecContext(
		ctx,
		query,
		referrerUserID, referalUserID,
	)
	if err != nil {
		return fmt.Errorf("can't save referal: %w", err)
	}

	return nil
}

func (r *Repository) SaveWithDate(ctx context.Context, referrerUserID int64, referalUserID int64, createdAt time.Time) error {
	query := fmt.Sprintf(`INSERT INTO %s.referals (referrer_user_id, referal_user_id,created_at) 
	VALUES ($1, $2,$3) `,
		r.db.Scheme())

	_, err := r.db.ExecContext(
		ctx,
		query,
		referrerUserID, referalUserID, createdAt,
	)
	if err != nil {
		return fmt.Errorf("can't save referal: %w", err)
	}

	return nil
}

func (r *Repository) GetAll(ctx context.Context) ([]telegram.Referal, error) {
	query := fmt.Sprintf(`SELECT referrer_user_id, referal_user_id, created_at 
FROM %s.referals`, r.db.Scheme())

	rows, err := r.db.QueryContext(
		ctx,
		query,
	)
	if err != nil {
		return nil, fmt.Errorf("can't get referals: %w", err)
	}

	defer rows.Close()

	var referals []telegram.Referal

	for rows.Next() {
		var referal telegram.Referal
		if err := rows.Scan(&referal.ReferrerID, &referal.ReferalID, &referal.CreatedAt); err != nil {
			return nil, fmt.Errorf("can't scan referal: %w", err)
		}

		referals = append(referals, referal)
	}

	return referals, nil
}

func (r *Repository) GetByReferrerID(ctx context.Context, referrerID uint64) ([]telegram.Referal, error) {
	query := fmt.Sprintf(`SELECT referrer_user_id, referal_user_id, created_at 
	FROM %s.referals WHERE referrer_user_id = $1`, r.db.Scheme())

	rows, err := r.db.QueryContext(ctx, query, referrerID)
	if err != nil {
		return nil, fmt.Errorf("can't get referals: %w", err)
	}

	defer rows.Close()

	var referals []telegram.Referal
	for rows.Next() {
		var referal telegram.Referal
		if err := rows.Scan(&referal.ReferrerID, &referal.ReferalID, &referal.CreatedAt); err != nil {
			return nil, fmt.Errorf("can't scan referal: %w", err)
		}

		referals = append(referals, referal)
	}

	return referals, nil
}
