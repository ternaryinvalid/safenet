package cache

import (
	"github.com/ternaryinvalid/safenet/client/internal/app/domain/config"
)

type Cache struct {
	config config.Cache
}

func New(config config.Cache) *Cache {
	return &Cache{
		config: config,
	}
}
