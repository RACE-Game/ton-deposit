package rest

import (
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"text/template"

	"github.com/RACE-Game/ton-deposit/interfaces"
)

// handleSomething обрабатывает веб-запрос
func handleSomething() http.Handler {
	// thing := prepareThing()
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// используем нечто для обработки запроса
			// log.Info(r.Context(), "msg", "handleSomething")
		},
	)
}

// server func
func healthzHandler(healthy *int32) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(healthy) == 1 {
			_ = encode(w, r, http.StatusOK, map[string]string{"status": "OK"})
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}
func readyzHandler(ready *int32) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(ready) == 1 {
			_ = encode(w, r, http.StatusOK, map[string]string{"status": "OK"})
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

func handleTemplate(files ...string) http.HandlerFunc {
	var (
		init   sync.Once
		tpl    *template.Template
		tplerr error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			tpl, tplerr = template.ParseFiles(files...)
		})
		if tplerr != nil {
			http.Error(w, tplerr.Error(), http.StatusInternalServerError)
			return
		}

		// используем tpl
		_ = tpl.Copy()
	}
}

func handleOK() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	)
}

func handleCORS() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
			w.Header().Set("Access-Control-Max-Age", "86400")

			w.WriteHeader(http.StatusNoContent)
		})
}

func HandlerReferalAll(us interfaces.UserService) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var referrerID uint64
			var refID *uint64
			referrerIDS := r.PathValue("referrer_id")

			if referrerIDS != "" {
				var err error
				referrerID, err = strconv.ParseUint(referrerIDS, 10, 64)
				if err != nil {
					_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid task id"})
					return
				}
				refID = &referrerID
			}

			referals, err := us.GetReferals(r.Context(), refID)
			if err != nil {
				_ = encode(w, r, http.StatusInternalServerError, map[string]string{"error": "can't get referals"})
				return
			}

			//scoresResponse := NewScoresResponse(scoresRequest.UserID, updatedScores)

			_ = encode(w, r, http.StatusOK, referals)
		},
	)
}

// func HandlerReferalIDCount(gc interfaces.GameController) http.HandlerFunc {
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, r *http.Request) {
// 			var referrerID uint64
// 			var refID *uint64
// 			referrerIDS := r.PathValue("referrer_id")

// 			if referrerIDS != "" {
// 				var err error
// 				referrerID, err = strconv.ParseUint(referrerIDS, 10, 64)
// 				if err != nil {
// 					_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid task id"})
// 					return
// 				}
// 				refID = &referrerID
// 			}

// 			referals, err := gc.GetReferalsCount(r.Context(), refID)
// 			if err != nil {
// 				_ = encode(w, r, http.StatusInternalServerError, map[string]string{"error": "can't get referals"})
// 				return
// 			}

// 			//scoresResponse := NewScoresResponse(scoresRequest.UserID, updatedScores)

// 			_ = encode(w, r, http.StatusOK, referals)
// 		},
// 	)
// }

// func HandlerSendInvite(tg interfaces.InviteSender) http.HandlerFunc {
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, r *http.Request) {

// 			userIDS := r.URL.Query().Get("id")
// 			if userIDS == "" {
// 				_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid id"})
// 				return
// 			}

// 			chatIDS := r.URL.Query().Get("chatid")
// 			if userIDS == "" {
// 				_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid user id"})
// 				return
// 			}

// 			chatID, err := strconv.ParseInt(chatIDS, 10, 64)
// 			if err != nil {
// 				_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid chat_id"})
// 			}

// 			userID, err := strconv.ParseInt(userIDS, 10, 64)
// 			if err != nil {
// 				_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid user_id"})
// 			}

// 			err = tg.SendInvite(r.Context(), userID, chatID)
// 			if err != nil {
// 				_ = encode(w, r, http.StatusInternalServerError, map[string]string{"error": "can't send invite"})
// 				return
// 			}

// 			//scoresResponse := NewScoresResponse(scoresRequest.UserID, updatedScores)

// 			_ = encode(w, r, http.StatusOK, map[string]string{"status": "OK"})
// 		},
// 	)
// }

// func HandlerNotify(nc interfaces.NotifierController) http.HandlerFunc {
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, r *http.Request) {
// 			value := r.Header.Get("Secure-Notify")
// 			if value != "5W9fLHnJGmWvjpHLGrTu" {
// 				_ = encode(w, r, http.StatusNotFound, map[string]string{"error": "not found"})
// 				return
// 			}

// 			userIDS := r.PathValue("user_id")

// 			var userID int64

// 			if userIDS != "" {
// 				userIDParsed, err := strconv.ParseInt(userIDS, 10, 64)
// 				if err != nil {
// 					_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid user id"})
// 				}
// 				userID = userIDParsed
// 			}

// 			notificationReq, err := decode[NotificationRequest](r)
// 			if err != nil {
// 				_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid notification"})
// 				return
// 			}

// 			n := notification.New(
// 				notificationReq.Text,
// 				notificationReq.Button.Text,
// 				notificationReq.Button.Type,
// 				notificationReq.Button.URL,
// 			)

// 			err = nc.SendNotification(r.Context(), userID, n)
// 			if err != nil {
// 				_ = encode(w, r, http.StatusInternalServerError, map[string]string{"error": "can't send notification"})
// 				return
// 			}

// 			_ = encode(w, r, http.StatusOK, map[string]string{"status": "OK"})
// 		},
// 	)
// }

// func HandlerNotifyAll(nc interfaces.NotifierController) http.HandlerFunc {
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, r *http.Request) {
// 			value := r.Header.Get("Secure-Notify")
// 			if value != "5W9fLHnJGmWvjpHLGrTu" {
// 				_ = encode(w, r, http.StatusNotFound, map[string]string{"error": "not found"})
// 				return
// 			}

// 			notificationReq, err := decode[NotificationRequest](r)
// 			if err != nil {
// 				_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid notification"})
// 				return
// 			}

// 			n := notification.New(
// 				notificationReq.Text,
// 				notificationReq.Button.Text,
// 				notificationReq.Button.Type,
// 				notificationReq.Button.URL,
// 			)

// 			err = nc.SendNotificationToAll(r.Context(), n)
// 			if err != nil {
// 				_ = encode(w, r, http.StatusInternalServerError, map[string]string{"error": "can't send notification"})
// 				return
// 			}

// 			_ = encode(w, r, http.StatusOK, map[string]string{"status": "OK"})
// 		},
// 	)
// }

func HandlerGroupMember(nc interfaces.NotifierController) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			userIDParam := r.URL.Query().Get("user_id")
			if userIDParam == "" {
				_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid user id"})
				return
			}

			chatIDParam := r.URL.Query().Get("chat_id")
			if chatIDParam == "" {
				_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid chat id"})
				return
			}

			userID, err := strconv.ParseInt(userIDParam, 10, 64)
			if err != nil {
				_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid user id"})
			}

			chatID, err := strconv.ParseInt(chatIDParam, 10, 64)
			if err != nil {
				_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid chat id"})
			}

			isMember, err := nc.IsMember(r.Context(), userID, chatID)
			if err != nil {
				_ = encode(w, r, http.StatusInternalServerError, map[string]string{"error": "can't get info about"})
				return
			}

			_ = encode(w, r, http.StatusOK, map[string]bool{"member": isMember})
		},
	)
}

func HandlerSaveUserData(us interfaces.UserService) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			userIDS := r.PathValue("user_id")

			var userID int64

			if userIDS != "" {
				userIDParsed, err := strconv.ParseInt(userIDS, 10, 64)
				if err != nil {
					_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid user id"})
				}
				userID = userIDParsed
			}

			userDataRequest, err := decode[UserDataRequest](r)
			if err != nil {
				_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid user data"})
				return
			}

			err = us.SaveUserData(r.Context(), userID, []byte(userDataRequest.Data))
			if err != nil {
				_ = encode(w, r, http.StatusInternalServerError, map[string]string{"error": "save user data"})
				return
			}

			_ = encode(w, r, http.StatusOK, map[string]string{"status": "OK"})
		},
	)
}

func HandlerGetUserData(uc interfaces.UserService) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			userIDS := r.PathValue("user_id")

			var userID int64

			if userIDS != "" {
				userIDParsed, err := strconv.ParseInt(userIDS, 10, 64)
				if err != nil {
					_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid user id"})
				}
				userID = userIDParsed
			}

			data, err := uc.GetUserData(r.Context(), userID)
			if err != nil {

				_ = encode(w, r, http.StatusInternalServerError, map[string]string{"error": "get user data"})
				return
			}

			_ = encode(w, r, http.StatusOK, map[string]string{"data": string(data)})
		},
	)
}

// func HandlerClaimTx(cc interfaces.ClaimController, logger interfaces.Logger) http.Handler {
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, r *http.Request) {
// 			txHash := r.PathValue("tx_hash")
// 			if txHash == "" {
// 				_ = encode(w, r, http.StatusBadRequest, map[string]string{"error": "invalid tx hash"})
// 				return
// 			}

// 			err := cc.CheckTx(r.Context(), txHash)
// 			if err != nil {
// 				logger.Error("msg", "can't confirm tx", err)
// 				_ = encode(w, r, http.StatusInternalServerError, map[string]string{"error": "invalid tx"})
// 				return
// 			}

// 			w.WriteHeader(http.StatusOK)
// 		},
// 	)
// }

// func HandlerGetClaims(cc interfaces.ClaimController, logger interfaces.Logger) http.Handler {
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, r *http.Request) {
// 			tokens, err := cc.GetClaims(r.Context())
// 			if err != nil {
// 				logger.Error("msg", "can't get tokens", err)
// 				_ = encode(w, r, http.StatusInternalServerError, map[string]string{"error": "invalid token"})
// 				return
// 			}

// 			_ = encode(w, r, http.StatusOK, tokens)
// 		},
// 	)
// }
