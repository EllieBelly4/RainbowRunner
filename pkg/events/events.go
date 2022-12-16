package events

import "reflect"

type IDHaver interface {
	ID() int
}

type EventHandlerEntry[T any] struct {
	Handler func(T)
	ID      *int
}

var globalSourceEventsMap = make(map[any][]interface{})

type Events struct {
}

func Emit[TEvent any](event TEvent) {
	if handlers, ok := globalSourceEventsMap[reflect.TypeOf(event)]; ok {
		for _, handlerEntry := range handlers {
			eventHandlerEntry := handlerEntry.(*EventHandlerEntry[TEvent])

			if eventHandlerEntry.ID != nil {
				continue
			}

			eventHandlerEntry.Handler(event)
		}
	}
}

func RegisterHandler[TEvent any](handler func(TEvent)) {
	eventType := reflect.TypeOf(*new(TEvent))
	handlers, ok := globalSourceEventsMap[eventType]

	if !ok {
		handlers = make([]interface{}, 0)
	}

	handlers = append(handlers, &EventHandlerEntry[TEvent]{
		Handler: handler,
	})

	globalSourceEventsMap[eventType] = handlers
}
