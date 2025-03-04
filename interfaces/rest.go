package interfaces

import (
	"context"

	"github.com/RACE-Game/ton-deposit-service/internal/domain/deposit"
	"github.com/RACE-Game/ton-deposit-service/internal/domain/notification"
	"github.com/RACE-Game/ton-deposit-service/internal/domain/telegram"
	"github.com/google/uuid"
)

type DepositService interface {
	CreateDeposit(ctx context.Context, userID int64, token string, amount uint64) error
	GetDeposits(ctx context.Context, userID int64) (deposites []deposit.Deposit, err error)
}

type GameService interface {
	GetReferals(ctx context.Context, referrerID *uint64) ([]telegram.Referal, error)
	GetReferalsCount(ctx context.Context, referrerID *uint64) (int64, error)
	//GetScores(ctx context.Context, limit *int64, userID *int64, token *string) ([]model.Score, error)
	//GetUserRank(ctx context.Context, userID int64, token string) (model.Rank, error)
}

type InviteSender interface {
	SendInvite(ctx context.Context, userID int64, chatID int64) error
}

type NotifierController interface {
	SendNotification(ctx context.Context, userID int64, notification notification.Notification) error
	IsMember(ctx context.Context, userID, chatID int64) (bool, error)
	SendNotificationToAll(ctx context.Context, notification notification.Notification) error
}

type UserService interface {
	//SaveWallet(ctx context.Context, wallet *model.Wallet) error
	SaveUserData(ctx context.Context, userID int64, data []byte) error
	GetUserData(ctx context.Context, userID int64) (data []byte, err error)
	Deposite(ctx context.Context, userID int64, token string, amount uint64) (uuid.UUID, error)
	GetAccount(ctx context.Context, userID int64) (telegram.Account, error)
	GetReferals(ctx context.Context, referrerID *uint64) ([]telegram.Referal, error)
}

// type ClaimController interface {
// 	Claim(ctx context.Context, token string, amount int64, wallet string) (claim *model.ClaimTxMeta, err error)
// 	CheckTx(ctx context.Context, txHash string) error
// 	GetClaims(ctx context.Context) ([]model.Token, error)
// }
