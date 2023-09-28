package repo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/cockroachdb/errors"
	"github.com/koki-develop/lgtmgen/backend/internal/env"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
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

type NotificationType string

const (
	NotificationTypeLGTMCreated NotificationType = "lgtm_created"
)

type NotificationMessage struct {
	Type        NotificationType    `json:"type"`
	LGTMCreated *LGTMCreatedMessage `json:"lgtm_created"`
}

type LGTMCreatedMessage struct {
	LGTM     *models.LGTM `json:"lgtm"`
	Source   string       `json:"source"`
	ClientIP string       `json:"client_ip"`
}

func (r *notificationsRepository) SendLGTMCreatedMessage(ctx context.Context, msg *LGTMCreatedMessage) error {
	b, err := json.Marshal(&NotificationMessage{
		Type:        NotificationTypeLGTMCreated,
		LGTMCreated: msg,
	})
	if err != nil {
		return errors.Wrap(err, "failed to marshal")
	}

	_, err = r.queueClient.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    util.Ptr(env.Vars.SQSQueueURLNotifications),
		MessageBody: util.Ptr(string(b)),
	})
	if err != nil {
		return errors.Wrap(err, "failed to send message")
	}

	return nil
}

func (r *notificationsRepository) NotifyLGTMCreated(ctx context.Context, msg *LGTMCreatedMessage) error {
	channel := r.channel()
	blocks := []slack.Block{
		slack.NewHeaderBlock(
			slack.NewTextBlockObject(slack.PlainTextType, "LGTM Created", false, false),
		),
		slack.NewSectionBlock(
			nil,
			[]*slack.TextBlockObject{
				slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("*ID*\n%s", msg.LGTM.ID), false, false),
				slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("*Source*\n%s", msg.Source), false, false),
				slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("*Client IP*\n%s", msg.ClientIP), false, false),
			},
			slack.NewAccessory(
				slack.NewImageBlockElement("https://koki.me/images/profile.png", "LGTM"), // TODO: set lgtm image url
			),
		),
	}
	log.Info(ctx, "notify lgtm created", "channel", channel, "blocks", blocks)

	_, _, err := r.slackClient.PostMessage(channel, slack.MsgOptionBlocks(blocks...))
	if err != nil {
		return errors.Wrap(err, "failed to post message")
	}

	return nil
}

func (r *notificationsRepository) channel() string {
	return fmt.Sprintf("#lgtmgen-%s", env.Vars.Stage)
}
