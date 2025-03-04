package user

import (
	"context"
	"fmt"
	"os/user"

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
func (r *Repository) Exist(ctx context.Context, userID int64) (exist bool, err error) {
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s.users WHERE telegram_id=$1)`, r.db.Scheme())
	rows, err := r.db.QueryContext(
		ctx,
		query,
		userID,
	)
	if err != nil {
		return false, fmt.Errorf("can't check user existance: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&exist); err != nil {
			return false, fmt.Errorf("can't scan user existance: %w", err)
		}
	}

	return exist, nil
}

func (r *Repository) GetByUserID(ctx context.Context, userID int64) (user.User, error) {
	// query := fmt.Sprintf("SELECT token_name, score  FROM %s.scores WHERE user_id = $1", r.db.Scheme())
	// rows, err := r.db.QueryContext(ctx, query, UserID)
	// if err != nil {
	// 	return nil, fmt.Errorf("can't get scores: %w", err)
	// }

	// defer rows.Close()

	// var scores model.Scores
	// for rows.Next() {
	// 	var score User
	// 	if err := rows.Scan(&score.Name, &score.Amount); err != nil {
	// 		return nil, fmt.Errorf("can't scan score: %w", err)
	// 	}

	// 	scores = append(scores, model.Score{
	// 		Name:   score.Name,
	// 		Amount: score.Amount,
	// 		UserID: score.UserID,
	// 	})
	// }

	return user.User{}, nil
}

func (r *Repository) Save(ctx context.Context, user telegram.User) error {
	query := fmt.Sprintf(`INSERT INTO %s.users
	 (telegram_id, tg_user_name, 
	 first_name, last_name,language_code, is_premium,chat_id,
	 start_message_id,
	 last_message_at
	 ) VALUES ($1, $2, $3,$4,$5,$6,$7,$8,$9)`,
		r.db.Scheme(),
	)

	_, err := r.db.ExecContext(
		ctx,
		query,
		user.TelegramID,
		user.TelegramUserName,
		user.FirstName,
		user.LastName,
		user.LanguageCode,
		user.IsPremium,
		user.ChatID,
		user.StartMessageID,
		user.LastMessageAt,
	)
	if err != nil {
		return fmt.Errorf("can't save user: %w", err)
	}

	return nil
}
