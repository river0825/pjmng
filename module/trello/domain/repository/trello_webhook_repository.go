package repository

import (
	"context"
	"river0825/cleanarchitecture/module/trello/domain/model"
)

type TrelloWebhookRepository interface {
	CreateTrelloWebhook(ctx context.Context, webhook *model.TrelloWebhook) error
}
