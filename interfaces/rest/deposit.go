package rest

import (
	"net/http"

	webapps "github.com/Fuchsoria/telegram-webapps"
	"github.com/RACE-Game/ton-deposit/interfaces"
)

type CreateDepositeRequest struct {
	Token  string
	Amount uint64
}

func HandlerDepositRequest(depositService interfaces.DepositService, logger Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userI := r.Context().Value(TelegramUserKey)
		telegramUser, ok := userI.(webapps.WebAppUser)
		if !ok {
			logger.Infof("failed get user from ctx")
		}

		var req CreateDepositeRequest

		req, err := decode[CreateDepositeRequest](r)
		if err != nil {
			_ = encode(w, r, http.StatusInternalServerError, map[string]string{"error": "decode error"})
			return
		}

		err = depositService.CreateDeposit(r.Context(), int64(telegramUser.ID), req.Token, req.Amount)
		if err != nil {
			_ = encode(w, r, http.StatusInternalServerError, map[string]string{"error": "no account"})
			return
		}

		// _ = encode(w, r, http.StatusOK, nil)
		w.WriteHeader(http.StatusOK)

	}
}

func HandlerGetDeposits(depositService interfaces.DepositService, logger Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userI := r.Context().Value(TelegramUserKey)
		telegramUser, ok := userI.(webapps.WebAppUser)
		if !ok {
			logger.Infof("failed get user from ctx")
		}

		deposits, err := depositService.GetDeposits(r.Context(), int64(telegramUser.ID))
		if err != nil {
			_ = encode(w, r, http.StatusInternalServerError, map[string]string{"error": "no account"})
			return
		}

		_ = encode(w, r, http.StatusOK, deposits)

	}
}
