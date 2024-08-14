package eventbus

import (
	"context"

	"river0825/cleanarchitecture/core/arch/domain/event"
)

type Handler func(ctx context.Context, event *event.Event) error
type IEventBus interface {
	Publish(event *event.Event) error
	PublishAll(event []*event.Event) error
	Subscribe(topic event.Topic, handler Handler) error
	SubscribeAll(handler Handler) error
}
