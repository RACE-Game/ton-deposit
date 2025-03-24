package tonscan

import (
	"context"
	"time"
)

type TonScan struct {
	tonscanService TonScanService
	ticker         *time.Ticker
	logger         Logger
}

func New(tonscanService TonScanService, ticker *time.Ticker, logger Logger) *TonScan {
	return &TonScan{
		tonscanService: tonscanService,
		ticker:         ticker,
		logger:         logger,
	}
}

func (ts *TonScan) Start(ctx context.Context) error {
	for {
		select {
		case <-ts.ticker.C:
			if err := ts.tonscanService.GetWalletData(ctx); err != nil {
				ts.logger.Errorf("tonscan error: %v", err)
			}
		case <-ctx.Done():
			return nil
		}
	}
}
