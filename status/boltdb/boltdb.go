package boltdb

import (
	"log"

	"time"

	"github.com/asdine/storm"
	"github.com/mirzakhany/pkg/status"
)

// New func implements the storage interface for status
func New(config status.ConfStatus) *KVStorage {
	return &KVStorage{
		config: config,
	}
}

// KVStorage is interface structure
type KVStorage struct {
	config status.ConfStatus
}

// Init client storage.
func (s *KVStorage) Init() error {
	return nil
}

// Reset Client storage.
func (s *KVStorage) Reset() {
}

func (s *KVStorage) setBoltDB(key string, value interface{}) error {
	db, err := storm.Open(s.config.BoltDB.Path)
	err = db.Set(s.config.BoltDB.Bucket, key, value)
	if err != nil {
		log.Println("BoltDB set error:", err.Error())
		return err
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Println("BoltDB error:", err.Error())
		}
	}()
	return err
}

func (s *KVStorage) delBoltDB(key string) error {
	db, err := storm.Open(s.config.BoltDB.Path)
	err = db.Delete(s.config.BoltDB.Bucket, key)
	if err != nil {
		log.Println("BoltDB delete error:", err.Error())
		return err
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println("BoltDB error:", err.Error())
		}
	}()
	return err
}

func (s *KVStorage) getBoltDB(key string, value interface{}) error {
	db, err := storm.Open(s.config.BoltDB.Path)
	err = db.Get(s.config.BoltDB.Bucket, key, value)
	if err != nil {
		log.Println("BoltDB get error:", err.Error())
		return err
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Println("BoltDB error:", err.Error())
		}
	}()
	return err
}

// IncIntKey inc value of an integer
func (s *KVStorage) IncIntKey(key string, value int64) error {
	val, err := s.GetIntKey(key)
	if err != nil {
		return err
	}
	total := val + value
	return s.setBoltDB(key, total)
}

// DecIntKey dec value of an integer
func (s *KVStorage) DecIntKey(key string, value int64) error {
	val, err := s.GetIntKey(key)
	if err != nil {
		return err
	}
	total := val + value
	return s.setBoltDB(key, total)
}

// SetIntKey set int value of a key
func (s *KVStorage) SetIntKey(key string, value int64, exp time.Duration) error {
	return s.setBoltDB(key, value)
}

// GetKey get int value of a key
func (s *KVStorage) GetIntKey(key string) (int64, error) {
	var value int64
	err := s.getBoltDB(key, &value)
	return value, err
}

// SetString key value of a key
func (s *KVStorage) SetString(key string, value string, exp time.Duration) error {
	return s.setBoltDB(key, value)
}

// GetString get value of a key
func (s *KVStorage) GetString(key string) (string, error) {
	var value string
	err := s.getBoltDB(key, &value)
	return value, err
}

// RemoveKeys remove keys
func (s *KVStorage) RemoveKeys(keys ...string) error {
	for _, key := range keys {
		err := s.delBoltDB(key)
		if err!= nil{
			return err
		}
	}
	return nil
}