package repo

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/cockroachdb/errors"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
	"github.com/koki-develop/lgtmgen/backend/internal/models"
	"github.com/koki-develop/lgtmgen/backend/internal/util"
	"github.com/slack-go/slack"
)

type notificationsRepository struct {
	slackClient *slack.Client
	queueClient *sqs.Client
}

func newNotificationsRepository(queueClient *sqs.Client, slackClient *slack.Client) *notificationsRepository {
	return &notificationsRepository{
		slackClient: slackClient,
		queueClient: queueClient,
	}
}

func (r *notificationsRepository) SendLGTMCreatedMessage(ctx context.Context, lgtm *models.LGTM) error {
	msg, err := json.Marshal(map[string]interface{}{
		"lgtm": lgtm,
	})
	if err != nil {
		return errors.Wrap(err, "failed to marshal")
	}

	_, err = r.queueClient.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    util.Ptr(env.Vars.SQSQueueURLNotifications),
		MessageBody: util.Ptr(string(msg)),
	})
	if err != nil {
		return errors.Wrap(err, "failed to send message")
	}

	return nil
}
