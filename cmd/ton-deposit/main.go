package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"

	"github.com/RACE-Game/ton-deposit/application"
	"github.com/RACE-Game/ton-deposit/infrastructure/db/deposit"
	"github.com/RACE-Game/ton-deposit/infrastructure/db/postgres"
	"github.com/RACE-Game/ton-deposit/interfaces/rest"
	"github.com/RACE-Game/ton-deposit/interfaces/ton"
)

type Config struct {
	HTTPHost           string `env:"HTTP_HOST" env-default:"127.0.0.1"`
	HTTPPort           string `env:"HTTP_PORT" env-default:"9001"`
	AppSecret          string `env:"APP_Secret" env-default:""`
	AppName            string `env:"APP_NAME" env-default:""`
	TonScanURL         string `env:"TON_SCAN_URL" env-default:"127.0.0.0:9003"`
	PostgresURL        string `env:"POSTGRES_URL" env-default:"" required:"true"`
	TelegramAPIKey     string `env:"TELEGRAM_API_KEY" env-default:""`
	TelegramBotLink    string `env:"TELEGRAM_BOT_LINK" env-default:""`
	MiniAppURL         string `env:"MINI_APP_URL" env-default:"https://cb7c912d.memezoo.pages.dev"`
	CommunityLink      string `env:"COMMUNITY_LINK" env-default:""`
	InviteMessage      string `env:"INVITE_MESSAGE" env-default:""`
	StartPicture       string `env:"START_PICTURE" env-default:""`
	TonWalletWordsSeed string `env:"TON_WALLET_WORDS_SEED" required:"false" env-default:""`
}

func run(ctx context.Context, _ io.Writer, args []string) error {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return fmt.Errorf("can't read config: %w", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	defer logger.Sync()

	sugarLogger := logger.Sugar()
	sugarLogger.Infof("Starting server with config: %s\n, %s\n",
		cfg.HTTPHost, cfg.HTTPPort,
	)

	tonClient := ton.New(cfg.TonScanURL)
	_ = tonClient

	db, err := postgres.New(ctx, cfg.PostgresURL, cfg.AppName, 4)
	if err != nil {
		return fmt.Errorf("can't create db: %w", err)
	}

	depositRepo, err := deposit.New(db)
	if err != nil {
		return fmt.Errorf("can't create deposit repository: %w", err)
	}

	err = depositRepo.Init(ctx)
	if err != nil {
		return fmt.Errorf("can't init deposit repository: %w", err)
	}

	depositService := application.NewDepositService(depositRepo, tonClient)

	restService := rest.NewServerMux(sugarLogger, cfg.TelegramAPIKey, cfg.AppSecret, depositService)

	server := rest.New(restService, cfg.HTTPHost, cfg.HTTPPort, sugarLogger)
	server.Start()
	// define http server
	// define service

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		sugarLogger.Info("Shutting down...")
		// only in proxy
	}()

	wg.Wait()

	return nil
}

func main() {
	ctx := context.Background()

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	defer cancel()

	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

/* docs
os.Args []string Аргументы, передаваемые при исполнении вашей программы. Также используется для флагов парсинга.
os.Stdin io.Reader Для считывания ввода
os.Stdout io.Writer Для записи вывода
os.Stderr io.Writer Для записи логов ошибок
os.Getenv func(string) string Для чтения переменных окружения
os.Getwd func() (string, error) Получение рабочей папки
*/
