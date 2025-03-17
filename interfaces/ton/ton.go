package ton

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/RACE-Game/ton-deposit/internal/domain/deposit"
)

func (c *Client) GetWallet(ctx context.Context, wallet string) (incomes []deposit.Deposit, err error) {
	path := c.tonScanURL + "/wallet/" + wallet

	resp, err := c.client.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("can't get wallet: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&incomes)
	if err != nil {
		return nil, fmt.Errorf("can't decode wallet: %w", err)
	}
	return incomes, nil
}
