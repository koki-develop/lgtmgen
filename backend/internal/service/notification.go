package service

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/koki-develop/lgtmgen/backend/internal/repo"
	"github.com/pkg/errors"
)

type notificationService struct {
	repo *repo.Repository
}

func newNotificationService(repo *repo.Repository) *notificationService {
	return &notificationService{
		repo: repo,
	}
}

func (s *notificationService) Notify(ctx context.Context, event *events.SQSEvent) error {
	for _, record := range event.Records {
		var msg repo.NotificationMessage
		if err := json.Unmarshal([]byte(record.Body), &msg); err != nil {
			return errors.Wrap(err, "failed to unmarshal")
		}

		var err error
		switch msg.Type {
		case repo.NotificationTypeLGTMCreated:
			err = s.repo.NotifyLGTMCreated(ctx, msg.LGTMCreated)
		case repo.NotificationTypeReportCreated:
			err = s.repo.NotifyReportCreated(ctx, msg.ReportCreated)
		default:
			err = errors.Errorf("unknown notification type: %s", msg.Type)
		}
		if err != nil {
			return errors.Wrap(err, "failed to notify")
		}
	}

	return nil
}
