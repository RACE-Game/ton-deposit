package application

import (
	"context"

	"github.com/RACE-Game/ton-deposit/internal/domain/deposit"
	"github.com/google/uuid"
)

type DepositRepository interface {
	Order(ctx context.Context, token string, userID int64, amount uint64, wallet string) (id uuid.UUID, err error)
	GetOrders(ctx context.Context) (orders []deposit.Order, err error)
}

type TonClient interface {
	GetWallet(ctx context.Context, wallet string) (incomes []deposit.Deposit, err error)
}

type Logger interface{}
