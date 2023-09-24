package service

type Service struct {
	*healthService
}

func New() *Service {
	return &Service{
		healthService: newHealthService(),
	}
}
