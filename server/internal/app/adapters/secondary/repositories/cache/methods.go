package cache

import (
	"errors"
	"log"
)

func (c *Cache) GetShared() (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.shared) == 0 {
		err := errors.New("не найден транспортный ключ")

		return "", err
	}

	return c.shared, nil
}

func (c *Cache) SaveShared(sharedSecret string) {
	c.mu.Lock()
	c.shared = sharedSecret
	c.mu.Unlock()

	log.Println("записан транспортный ключ в кэш")
}

func (c *Cache) SetSecret(localPublicKey, localPrivateKey string) {
	c.mu.Lock()
	c.pubKey = localPublicKey
	c.privateKey = localPrivateKey
	c.mu.Unlock()
}

func (c *Cache) GetSecret() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.privateKey
}

func (c *Cache) IsEmpty() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.pubKey) == 0 || len(c.privateKey) == 0 {
		return true
	}

	return false
}
