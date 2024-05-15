package cache

import "sync"

type Cache struct {
	mu         sync.Mutex
	shared     string
	pubKey     string
	privateKey string
}

func New() *Cache {
	return &Cache{
		mu:         sync.Mutex{},
		shared:     "",
		pubKey:     "",
		privateKey: "",
	}
}
