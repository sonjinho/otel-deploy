package services

import (
	"log"

	"github.com/dgraph-io/badger/v4"
	"github.com/sonjinho/otel-deploy-backend/internal/database"
)

type ConfigService struct {
	DB *database.BadgerDB
}

func NewConfigService(db *database.BadgerDB) *ConfigService {
	return &ConfigService{DB: db}
}

func (s *ConfigService) GetConfig(key string) (string, error) {
	var value string
	err := s.DB.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			value = string(val)
			return nil
		})
	})
	return value, err
}

func (s *ConfigService) SaveConfig(key, value string) error {
	err := s.DB.DB.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
	if err == nil {
		log.Printf("Config saved: %s -> %s", key, value)
	}
	return err
}
