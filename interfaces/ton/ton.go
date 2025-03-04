package ton

import (
	"context"

	"github.com/RACE-Game/ton-deposit-service/internal/domain/deposit"
)

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c *Client) GetWallet(ctx context.Context, wallet string) (incomes []deposit.Deposit, err error) {
	return nil, nil
}
