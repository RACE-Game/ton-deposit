package deposit

import (
	"context"
	"database/sql"
	"fmt"
)

var ErrTokenNotFound = fmt.Errorf("claim token info not found")

func (r *Repository) GetToken(ctx context.Context, token string) (model.ClaimTokenInfo, error) {
	var t model.ClaimTokenInfo
	q := fmt.Sprintf(`SELECT 
			name,
			address,
			decimals,
			multiplicator,
			budget,
			meta,
			active,
			custom_score_table,
			created_at,
			updated_at
		FROM
			%s.tokens
		WHERE
			name = $1 and active = true`, r.db.Scheme())
	err := r.db.QueryRowContext(ctx, q, token).Scan(
		&t.Name,
		&t.Address,
		&t.Decimals,
		&t.Multiplicator,
		&t.Budget,
		&t.Data,
		&t.Active,
		&t.TableName,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.ClaimTokenInfo{}, ErrTokenNotFound
		}
		return model.ClaimTokenInfo{}, fmt.Errorf("can't get token: %w", err)
	}
	return t, nil
}

func (r *Repository) GetTokenAll(ctx context.Context) (map[string]model.ClaimTokenInfo, error) {
	var t model.ClaimTokenInfo
	q := fmt.Sprintf(`SELECT 
			name,
			address,
			decimals,
			multiplicator,
			budget,
			meta,
			active,
			custom_score_table,
			created_at,
			updated_at
		FROM
			%s.tokens
			WHERE active = true`,
		r.db.Scheme())
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	return nil, ErrTokenNotFound
		// }
		return nil, fmt.Errorf("can't get token: %w", err)
	}

	defer rows.Close()

	tokens := make(map[string]model.ClaimTokenInfo)

	for rows.Next() {
		err := rows.Scan(
			&t.Name,
			&t.Address,
			&t.Decimals,
			&t.Multiplicator,
			&t.Budget,
			&t.Data,
			&t.Active,
			&t.TableName,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("can't get token: %w", err)
		}

		tokens[t.Name] = t
	}

	return tokens, nil
}
