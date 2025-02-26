package rest

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/RACE-Game/ton-deposit-service/interfaces"
)

var (
	healthy int32
	ready   int32
	// watcher *fscache.Watcher
)

type Logger interface {
	Info(...interface{})
	Error(...interface{})
	Infof(string, ...interface{})
}
type Server struct {
	*http.Server

	logger Logger
}

func NewServerMux(
	logger Logger,
	tgAPIKey string,
	appSecret string,
	userService interfaces.UserService,
	depositService interfaces.DepositService,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		logger,
		userService,
		depositService,
	)

	// add app version to response
	handlerWithVersion := versionMiddleware(mux)
	//handlerShowHeader := showHeadersMiddleware(handlerWithVersion)
	handlerAuth := checkTelegramAuthMiddleware(handlerWithVersion, tgAPIKey, appSecret)

	return handlerAuth
}

func New(mux http.Handler, host, port string, logger Logger) *Server {
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(host /* config.Host */, port /* config.Port */),
		Handler: mux,
	}

	return &Server{
		Server: httpServer,
		logger: logger,
	}

}

func (s *Server) Start() {
	go func() {
		s.logger.Infof("listening on %s\n", s.Addr)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			atomic.StoreInt32(&healthy, 0)
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}

	}()

	atomic.StoreInt32(&healthy, 1)
}

func (s *Server) ListenAndServe() error {
	return s.Server.ListenAndServe()
}
