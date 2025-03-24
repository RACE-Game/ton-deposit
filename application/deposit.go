package application

import (
	"context"

	"github.com/RACE-Game/ton-deposit/internal/domain/deposit"
	"github.com/google/uuid"
)

type DepositService struct {
	depositRepo DepositRepository
	tonClient   TonClient
}

func NewDepositService(
	depositRepo DepositRepository,
	tonClient TonClient,
) *DepositService {
	return &DepositService{
		depositRepo: depositRepo,
		tonClient:   tonClient,
	}
}

func (s *DepositService) CreateOrder(ctx context.Context, userID int64, token, wallet string, amount uint64) (uuid.UUID, error) {
	orderID, err := s.depositRepo.Order(ctx, token, userID, amount, wallet)
	if err != nil {
		return uuid.Nil, err
	}

	return orderID, nil
}

func (s *DepositService) CreateDeposit(ctx context.Context, userID int64, token string, amount uint64) error {
	return nil
}
func (s *DepositService) GetDeposits(ctx context.Context, userID int64) (deposites []deposit.Deposit, err error) {
	return nil, nil
}
