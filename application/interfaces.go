package application

import (
	"context"

	"github.com/google/uuid"
)

type DepositRepository interface {
	Order(ctx context.Context, token string, userID int64, amount uint64, wallet string) (id uuid.UUID, err error)
}

type TonClient interface{}

type Logger interface{}
