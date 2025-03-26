package tonscan

import (
	"context"
)

type TonScanService interface {
	GetWalletData(ctx context.Context) (err error)
}

type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}
