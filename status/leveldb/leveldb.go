package leveldb

import (
	"fmt"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"

	"time"

	"github.com/mirzakhany/pkg/logger"
	"github.com/mirzakhany/pkg/status"
)

var dbPath string

func setLevelDB(key string, value interface{}) error {
	db, _ := leveldb.OpenFile(dbPath, nil)
	val := fmt.Sprintf("%d", value)

	err := db.Put([]byte(key), []byte(val), nil)

	defer func() {
		err := db.Close()
		if err != nil {
			logger.Error("LevelDB error:", err.Error())
		}
	}()
	return err
}

func getLevelDB(key string, value interface{}) error {
	db, err := leveldb.OpenFile(dbPath, nil)

	data, err := db.Get([]byte(key), nil)
	value, err = strconv.ParseInt(string(data), 10, 64)

	defer func() {
		err := db.Close()
		if err != nil {
			logger.Error("LevelDB error:", err.Error())
		}
	}()
	return err
}

func delLevelDB(key string) error {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return err
	}
	err = db.Delete([]byte(key), nil)
	defer func() {
		err := db.Close()
		if err != nil {
			logger.Error("LevelDB error:", err.Error())
		}
	}()
	return err
}

// New func implements the storage interface
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
	dbPath = s.config.LevelDB.Path
	return nil
}

// Reset Client storage.
func (s *KVStorage) Reset() {
}

// IncIntKey inc value of an integer
func (s *KVStorage) IncIntKey(key string, value int64) error {
	val, err := s.GetIntKey(key)
	if err != nil {
		return err
	}
	total := val + value
	return setLevelDB(key, total)
}

// DecIntKey dec value of an integer
func (s *KVStorage) DecIntKey(key string, value int64) error {
	val, err := s.GetIntKey(key)
	if err != nil {
		return err
	}
	total := val - value
	return setLevelDB(key, total)
}

// SetIntKey set int value of a key
func (s *KVStorage) SetIntKey(key string, value int64, exp time.Duration) error {
	return setLevelDB(key, value)
}

// GetKey get int value of a key
func (s *KVStorage) GetIntKey(key string) (int64, error) {
	var value int64
	err := getLevelDB(key, &value)
	return value, err
}

// SetString key value of a key
func (s *KVStorage) SetString(key string, value string, exp time.Duration) error {
	return setLevelDB(key, value)
}

// GetString get value of a key
func (s *KVStorage) GetString(key string) (string, error) {
	var value string
	err := getLevelDB(key, &value)
	return value, err
}

// RemoveKeys remove keys
func (s *KVStorage) RemoveKeys(keys ...string) error {
	for _, key := range keys {
		err := delLevelDB(key)
		if err != nil {
			return err
		}
	}
	return nil
}
