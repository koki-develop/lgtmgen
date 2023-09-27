package service

import (
	"context"

	"github.com/koki-develop/lgtmgen/backend/internal/log"
)

type notificationService struct{}

func newNotificationService() *notificationService {
	return &notificationService{}
}

func (s *notificationService) NotifyLGTMCreated(ctx context.Context, ipt interface{}) error {
	// TODO: implement
	log.Info(ctx, "notify", "ipt", ipt)
	return nil
}
