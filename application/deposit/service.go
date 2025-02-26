package deposit

import "context"

type Service struct {
	tonClient         TonClient
	depositRepository DepositRepository
	logger            Logger
}

func New(
	tonClient TonClient,
	depositRepository DepositRepository,
	logger Logger,
) *Service {
	return &Service{
		tonClient:         tonClient,
		depositRepository: depositRepository,
		logger:            logger,
	}
}

func (s *Service) Start(ctx context.Context) error {

	return nil
}
