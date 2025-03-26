package application

import (
	"context"
	"strings"
)

type TonScanService struct {
	depositRepository DepositRepository
	tonClient         TonClient
	wallet            string
}

func NewTonScanService(
	depositRepository DepositRepository,
	tonClient TonClient,
	wallet string,
) *TonScanService {
	return &TonScanService{
		depositRepository: depositRepository,
		tonClient:         tonClient,
		wallet:            wallet,
	}
}

func (ts *TonScanService) GetWalletData(ctx context.Context) error {
	orders, err := ts.depositRepository.GetOrders(ctx)
	if err != nil {
		return err
	}

	incomes, err := ts.tonClient.GetWallet(ctx, ts.wallet)
	if err != nil {
		return err
	}

	for _, income := range incomes {
		for _, order := range orders {
			if strings.Contains(income.Comment, order.ID.String()) {
				if err := ts.depositRepository.UpdateOrder(ctx, order.ID, income.TXHash); err != nil {
					return err
				}

				if err := ts.depositRepository.CreateDeposit(
					ctx,
					order.ID,
					order.UserID,
					order.Wallet,
					order.Token,
					income.Amount,
				); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
