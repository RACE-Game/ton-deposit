package rest

import (
	"net/http"

	webapps "github.com/Fuchsoria/telegram-webapps"
	"github.com/RACE-Game/ton-deposit-service/interfaces"
)

func HandlerGetAccount(userService interfaces.UserService, logger Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userI := r.Context().Value(TelegramUserKey)
		telegramUser, ok := userI.(webapps.WebAppUser)
		if !ok {
			logger.Infof("failed get user from ctx")
		}

		account, err := userService.GetAccount(r.Context(), int64(telegramUser.ID))
		if err != nil {
			_ = encode(w, r, http.StatusInternalServerError, map[string]string{"error": "no account"})

		}

		_ = encode(w, r, http.StatusOK, account)

	}
}
