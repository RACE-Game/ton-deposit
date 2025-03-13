package deposit

import (
	"fmt"

	"github.com/RACE-Game/ton-deposit/infrastructure/db"
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
