package geventbus

import "sync"

type EventBus struct {
	subscribers map[string][]func(interface{})
	lock        sync.RWMutex
}

var (
	instance *EventBus
	once     sync.Once
)

func GetInstance() *EventBus {
	once.Do(func() {
		instance = &EventBus{
			subscribers: make(map[string][]func(interface{})),
		}
	})
	return instance
}

// Subscribe 订阅
func (bus *EventBus) Subscribe(eventType string, handler func(interface{})) {
	bus.lock.Lock()
	defer bus.lock.Unlock()
	bus.subscribers[eventType] = append(bus.subscribers[eventType], handler)
}

// Publish 触发
func (bus *EventBus) Publish(eventType string, data interface{}) {
	bus.lock.RLock()
	defer bus.lock.RUnlock()
	if handlers, found := bus.subscribers[eventType]; found {
		for _, handler := range handlers {
			handler(data)
		}
	}
}
