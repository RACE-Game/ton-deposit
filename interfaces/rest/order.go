package rest

import (
	"net/http"

	webapps "github.com/Fuchsoria/telegram-webapps"
	"github.com/RACE-Game/ton-deposit/interfaces"
)

func HandlerCreateOrder(orderService interfaces.DepositService, logger Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userI := r.Context().Value(TelegramUserKey)
		telegramUser, ok := userI.(webapps.WebAppUser)
		if !ok {
			logger.Infof("failed get user from ctx")
		}

		var req CreateOrderRequest

		req, err := decode[CreateOrderRequest](r)
		if err != nil {
			_ = encode(
				w,
				r,
				http.StatusInternalServerError,
				map[string]string{"error": "decode error"},
			)
			return
		}

		orderID, err := orderService.CreateOrder(r.Context(), int64(telegramUser.ID), req.Token, req.Wallet, req.Amount)
		if err != nil {
			_ = encode(
				w,
				r,
				http.StatusInternalServerError,
				map[string]string{"error": "no account"},
			)
			return
		}

		_ = encode(w, r, http.StatusOK, CreateOrderResponse{OrderID: orderID})
		w.WriteHeader(http.StatusOK)
	}
}
