package cache

import (
	"encoding/json"
	"github.com/ternaryinvalid/safenet/client/internal/app/domain/entity"
	"log"
	"os"
)

func (c *Cache) SaveAccount(account *entity.Account) error {
	log.Println(os.Getwd())
	log.Println(c.config.Filepath)

	jsonData, err := json.Marshal(account)
	if err != nil {
		return err
	}

	err = os.WriteFile(c.config.Filepath, jsonData, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache) LoadAccount() (*entity.Account, error) {
	jsonData, err := os.ReadFile(c.config.Filepath)
	if err != nil {
		return nil, err
	}

	var account entity.Account
	err = json.Unmarshal(jsonData, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
