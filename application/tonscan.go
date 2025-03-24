package application

import "context"

type TonScanService struct {
	depositRepository DepositRepository
	tonClient         TonClient
}

func NewTonScanService(
	depositRepository DepositRepository,
	tonClient TonClient,
) *TonScanService {
	return &TonScanService{
		depositRepository: depositRepository,
		tonClient:         tonClient,
	}
}

func (ts *TonScanService) TonScan(ctx context.Context) error {
	ts.tonClient.GetWallet(ctx context.Context, wallet string) (incomes []deposit.Deposit, err error)

	return nil
}
