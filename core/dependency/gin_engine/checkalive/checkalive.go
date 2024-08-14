package checkalive

import "context"

type ICheckAlive interface {
	Ping(ctx context.Context) error
	Name() string
}
