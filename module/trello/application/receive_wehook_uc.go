package application

import (
	"context"

	"river0825/cleanarchitecture/core/arch/application"
	"river0825/cleanarchitecture/module/trello/domain/model"
	"river0825/cleanarchitecture/module/trello/domain/repository"

	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
)

type WebhookCommand struct {
	Webhook     *model.TrelloWebhook
	HeaderMac   string
	Body        string
	CallbackURL string
	Secret      string
}

func (cmd *WebhookCommand) Validate() error {
	if cmd.Webhook == nil {
		return errors.New("webhook is required")
	}
	if cmd.HeaderMac == "" {
		return errors.New("header mac is required")
	}
	if cmd.Body == "" {
		return errors.New("body is required")
	}
	if cmd.CallbackURL == "" {
		return errors.New("callback url is required")
	}

	content := cmd.Body + cmd.CallbackURL
	mac := hmac.New(sha1.New, []byte(cmd.Secret))
	mac.Write([]byte(content))
	expectedMAC := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	if expectedMAC != cmd.HeaderMac {
		return errors.New("invalid webhook signature")
	}

	return nil
}

type webhookResponse struct {
	Success bool
	Message string
}

type webhookUseCase struct {
	repo repository.TrelloWebhookRepository
}

var _ application.IUseCase = (*webhookUseCase)(nil)

func (A *webhookUseCase) Execute(ctx context.Context, c any) (any, error) {
	cmd := c.(*WebhookCommand)
	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	// Example logic to handle the webhook
	err := A.repo.CreateTrelloWebhook(ctx, cmd.Webhook)
	if err != nil {
		return nil, err
	}

	rsp := &webhookResponse{
		Success: true,
		Message: "Webhook processed successfully",
	}

	return rsp, nil
}
