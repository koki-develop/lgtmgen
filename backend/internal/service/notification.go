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

type notificationType string

const (
	notificationTypeLGTMCreated notificationType = "lgtm_created"
)

type notificationMessage struct {
	Type        notificationType         `json:"type"`
	LGTMCreated *repo.LGTMCreatedMessage `json:"lgtm_created"`
}

func (s *notificationService) Notify(ctx context.Context, event *events.SQSEvent) error {
	for _, record := range event.Records {
		var msg notificationMessage
		if err := json.Unmarshal([]byte(record.Body), &msg); err != nil {
			return errors.Wrap(err, "failed to unmarshal")
		}

		var err error
		switch msg.Type {
		case notificationTypeLGTMCreated:
			err = s.notifyLGTMCreated(ctx, msg.LGTMCreated)
		}
		if err != nil {
			return errors.Wrap(err, "failed to notify")
		}
	}

	return nil
}

func (s *notificationService) notifyLGTMCreated(ctx context.Context, msg *repo.LGTMCreatedMessage) error {
	if err := s.repo.NotifyLGTMCreated(ctx, msg); err != nil {
		return errors.Wrap(err, "failed to notify lgtm created")
	}

	return nil
}
