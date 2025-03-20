package telegram

import (
	"context"

	"github.com/RACE-Game/ton-deposit/internal/domain/telegram"
)

type UserController interface {
	SaveUser(ctx context.Context, user telegram.User) error
	UserExist(ctx context.Context, userID int64) (bool, error)
	UpdateUser(ctx context.Context, user telegram.User) error
	GetReferalLink(ctx context.Context, userID int64) string
	GetGameLink(ctx context.Context) string
	UserJoined(ctx context.Context, referalOwnerUserID int64, referalUserID int64) error
}

type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}
