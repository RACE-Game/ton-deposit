package ton

import (
	"context"

	"github.com/RACE-Game/ton-deposit/internal/domain/deposit"
)

type Client struct {
	tonScanURL string
}

func New(tonScanURL string) *Client {
	return &Client{
		tonScanURL: tonScanURL,
	}
}

func (c *Client) GetWallet(ctx context.Context, wallet string) (incomes []deposit.Deposit, err error) {

	return nil, nil
}
