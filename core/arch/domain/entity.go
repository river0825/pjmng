package domain

import "river0825/cleanarchitecture/core/arch/domain/event"

type Entity struct {
	Events []*event.Event `json:"-"`
}

func (e *Entity) AddEvent(event *event.Event) {
	e.Events = append(e.Events, event)
}
