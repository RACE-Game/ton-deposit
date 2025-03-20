package user

import (
	"context"
	"fmt"
	"time"
)

func (c *Repository) GetUserData(ctx context.Context, userID int64) (data []byte, err error) {
	query := fmt.Sprintf(`SELECT data FROM %s.user_data WHERE user_id = $1`,
		c.db.Scheme(),
	)

	rows, err := c.db.QueryContext(
		ctx,
		query,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("can't get user data: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&data); err != nil {
			return nil, fmt.Errorf("can't scan user data: %w", err)
		}
	}

	return data, nil
}

func (r *Repository) SaveUserData(ctx context.Context, userID int64, data []byte) error {
	query := fmt.Sprintf(`INSERT INTO %s.user_data
	 (user_id, data,updated_at) VALUES ($1, $2,$3)`,
		r.db.Scheme(),
	)

	_, err := r.db.ExecContext(
		ctx,
		query,
		userID,
		data,
		time.Now().UTC(),
	)
	if err != nil {
		return fmt.Errorf("can't save user data: %w", err)
	}

	return nil
}

func (r *Repository) UpdateUserData(ctx context.Context, userID int64, data []byte) error {
	query := fmt.Sprintf(`UPDATE %s.user_data SET data = $1, updated_at = $2 WHERE user_id = $3`,
		r.db.Scheme(),
	)

	_, err := r.db.ExecContext(
		ctx,
		query,
		data,
		time.Now().UTC(),
		userID,
	)
	if err != nil {
		return fmt.Errorf("can't update user data: %w", err)
	}

	return nil
}

func (r *Repository) UserDataExists(ctx context.Context, userID int64) (bool, error) {
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s.user_data WHERE user_id = $1)`,
		r.db.Scheme(),
	)

	var exists bool

	rows, err := r.db.QueryContext(
		ctx,
		query,
		userID,
	)
	if err != nil {
		return false, fmt.Errorf("can't check user data existance: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&exists); err != nil {
			return false, fmt.Errorf("can't scan user data existance: %w", err)
		}
	}

	return exists, nil
}
