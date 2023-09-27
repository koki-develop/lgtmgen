package service

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/koki-develop/lgtmgen/backend/internal/models"
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

type lgtmMessageBody struct {
	LGTM *models.LGTM `json:"lgtm"`
}

func (s *notificationService) Notify(ctx context.Context, event *events.SQSEvent) error {
	var lgtms models.LGTMs
	for _, record := range event.Records {
		var msg lgtmMessageBody
		if err := json.Unmarshal([]byte(record.Body), &msg); err != nil {
			return err
		}
		lgtms = append(lgtms, msg.LGTM)
	}

	// TODO: notify to slack
	log.Info(ctx, "created lgtms", "lgtms", lgtms)

	return nil
}
