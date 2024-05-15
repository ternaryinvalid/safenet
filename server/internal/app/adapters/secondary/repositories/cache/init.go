package cache

import "sync"

type Cache struct {
	mu         sync.Mutex
	data       map[string]string // хранит shared key по public key ключу
	pubKey     string
	privateKey string
}

func New() *Cache {
	return &Cache{
		mu:         sync.Mutex{},
		data:       make(map[string]string),
		pubKey:     "",
		privateKey: "",
	}
}
