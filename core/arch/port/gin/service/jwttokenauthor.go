package service

import (
	"context"
	"river0825/cleanarchitecture/core/arch/domain/model"
)

type JWTTokenAuthor interface {
	ValidateAccessToken(ctx context.Context, token string) (*model.User, error)
}
