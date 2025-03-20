package rest

import (
	"net/http"

	"github.com/RACE-Game/ton-deposit/interfaces"
)

/*
POST /wallet/investor // setup investor wallet
GET /wallet/create/{n} // create N wallets
GET /wallet/count // count wallets
GET /wallet // list wallets
GET /wallet/{id} // get wallet info
POST /wallet/{id}/tags // edit tags
GET /tags // list all tags
GET /wallet/investor // investor wallet info
POST /wallet/invest // invest money from investor wallet to all
GET  /wallet/collect // collect all money to investor wallet
*/
func addRoutes(
	mux *http.ServeMux,
	logger Logger,
	userService interfaces.UserService,
	depositService interfaces.DepositService,
	orderService interfaces.OrderService,
) {
	mux.Handle("GET /healthz", healthzHandler(&healthy))
	mux.Handle("GET /readyz", readyzHandler(&ready))
	mux.Handle("GET /ping", healthzHandler(&healthy))
	mux.Handle("GET /", http.NotFoundHandler())

	mux.Handle("GET /referals", corsMiddleware(HandlerReferalAll(userService)))
	mux.Handle("GET /referals/{referrer_id}", corsMiddleware(HandlerReferalAll(userService)))
	// mux.Handle("GET /referals/{referrer_id}/count", corsMiddleware(HandlerReferalIDCount(userService)))
	mux.Handle("OPTIONS /referals/{referrer_id...}", corsMiddleware(handleCORS()))
	mux.Handle("OPTIONS /referals", corsMiddleware(handleCORS()))

	mux.Handle("GET /account", corsMiddleware(HandlerGetAccount(userService, logger)))
	mux.Handle("OPTIONS /account", corsMiddleware(handleCORS()))

	mux.Handle("POST /deposit", corsMiddleware(HandlerDepositRequest(depositService, logger)))
	mux.Handle("GET /deposit", corsMiddleware(HandlerGetDeposits(depositService, logger)))

	mux.Handle("OPTIONS /deposit", corsMiddleware(handleCORS()))

	mux.Handle("POST /order", corsMiddleware(HandlerCreateOrder(orderService, logger)))
	mux.Handle("OPTIONS /order", corsMiddleware(handleCORS()))
	// mux.Handle("GET /group/member", corsMiddleware(HandlerGroupMember(notifyController)))
	// mux.Handle("OPTIONS /group/member", corsMiddleware(handleCORS()))

	// (mux.Handle("GET /scores/okx", corsMiddleware(okxListHandler(okxController, logger))))
	// mux.Handle("GET /scores", corsMiddleware(handlers.HandlerGetScores(gameController)))
	// mux.Handle("GET /scores/{user_id}/{token}", corsMiddleware(handlers.HandlerGetPosition(gameController)))
	// mux.Handle("OPTIONS /scores", corsMiddleware(handlers.handleCORS()))
	// mux.Handle("OPTIONS /scores/{path...}", corsMiddleware(handlers.handleCORS()))

	// //mux.Handle("POST /notify", corsMiddleware(HandlerGetPosition(gameController)))
	// mux.Handle("POST /notify/all", corsMiddleware(handlers.HandlerNotifyAll(notifyController)))
	// mux.Handle("POST /notify/{user_id}", corsMiddleware(handlers.HandlerNotify(notifyController)))
	// //mux.Handle("POST /notify/all", corsMiddleware(HandlerNotify(notifyController)))

	// mux.Handle("OPTIONS /{p...}", corsMiddleware(handlers.handleCORS()))

	// mux.Handle("GET /invite", corsMiddleware(handlers.HandlerSendInvite(telegram)))

	// mux.Handle("OPTIONS /user/{user_id}/wallet", corsMiddleware(handlers.handleCORS()))
	// mux.Handle("POST /user/{user_id}/wallet", corsMiddleware(handlers.HandlerUserWallet(userController)))

	// mux.Handle("OPTIONS /user/{user_id}/data", corsMiddleware(handlers.handleCORS()))
	// mux.Handle("POST /user/{user_id}/data", corsMiddleware(handlers.HandlerSaveUserData(userController)))
	// mux.Handle("GET /user/{user_id}/data", corsMiddleware(handlers.HandlerGetUserData(userController)))

	// mux.Handle("OPTIONS /user/{user_id}/okx", corsMiddleware(handlers.handleCORS()))
	// mux.Handle("GET /user/{user_id}/okx", corsMiddleware(handlers.okxCheckHandler(okxController, logger)))

	// mux.Handle("POST /claim", corsMiddleware(handlers.HandlerClaim(claimController, logger)))
	// mux.Handle("OPTIONS /claim", corsMiddleware(handlers.handleCORS()))

	// mux.Handle("GET /claim", corsMiddleware(handlers.HandlerGetClaims(claimController, logger)))
	// mux.Handle("POST /claim/{tx_hash}", corsMiddleware(handlers.HandlerClaimTx(claimController, logger)))
	// mux.Handle("OPTIONS /claim/{p...}", corsMiddleware(handlers.handleCORS()))

	logger.Info("routes added")
}
