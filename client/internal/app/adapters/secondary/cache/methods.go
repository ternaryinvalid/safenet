package cache

import (
	"errors"
	"log"
)

func (c *Cache) GetShared(publicKey string) (string, error) {
	c.mu.Lock()
	shared, ok := c.data[publicKey]
	c.mu.Unlock()
	if !ok {
		err := errors.New("не найден транспортный ключ для такого публичного ключа")

		return "", err
	}

	return shared, nil
}

func (c *Cache) SaveShared(remotePublicKey, sharedSecret string) {
	c.mu.Lock()
	c.data[sharedSecret] = remotePublicKey
	c.mu.Unlock()

	log.Println("записана новая сессия c хостом")
}

func (c *Cache) SetSecret(localPublicKey, localPrivateKey string) {
	c.mu.Lock()
	c.pubKey = localPublicKey
	c.privateKey = localPrivateKey
	c.mu.Unlock()
}

func (c *Cache) Private() string {
	return c.privateKey
}

func (c *Cache) Public() string {
	return c.pubKey
}

func (c *Cache) IsEmpty() bool {
	if len(c.pubKey) == 0 || len(c.privateKey) == 0 {
		return true
	}

	return false
}
