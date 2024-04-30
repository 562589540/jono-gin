package cache

import (
	"sync"
)

var (
	instance *Service
	once     sync.Once
)

type Service struct {
	PathToRole map[string]uint
}

func New() *Service {
	once.Do(func() {
		instance = &Service{
			PathToRole: make(map[string]uint),
		}
	})
	return instance
}
