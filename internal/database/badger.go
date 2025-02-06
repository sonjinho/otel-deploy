package database

import (
	"log"

	"github.com/dgraph-io/badger/v4"
)

type BadgerDB struct {
	DB *badger.DB
}

func NewBadgerDB(dbPath string) (*BadgerDB, error) {
	opts := badger.DefaultOptions(dbPath)
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return &BadgerDB{DB: db}, nil
}

func (b *BadgerDB) Close() {
	if err := b.DB.Close(); err != nil {
		log.Println("Error closing BadgerDB:", err)
	}
}
