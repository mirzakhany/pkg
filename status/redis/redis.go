package redis

import (
	"strconv"

	"time"

	"github.com/go-redis/redis"
	"github.com/mirzakhany/pkg/logger"
	"github.com/mirzakhany/pkg/status"
)

//
var redisClient *redis.Client

// New func implements the storage interface
func New(config status.ConfStatus) *KVStorage {
	return &KVStorage{
		config: config,
	}
}

func getInt64(key string, value *int64) (err error) {
	val, _ := redisClient.Get(key).Result()
	*value, err = strconv.ParseInt(val, 10, 64)
	return err
}

func getString(key string, value *string) (err error) {
	*value, err = redisClient.Get(key).Result()
	return err
}

// KVStorage is interface structure
type KVStorage struct {
	config status.ConfStatus
}

// Init client storage.
func (s *KVStorage) Init() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     s.config.Redis.Addr,
		Password: s.config.Redis.Password,
		DB:       s.config.Redis.DB,
	})

	_, err := redisClient.Ping().Result()

	if err != nil {
		// redis server error
		logger.Fatalf("Can't connect redis server: " + err.Error())
		return err
	}

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
	return redisClient.Set(key, strconv.Itoa(int(total)), 0).Err()
}

// DecIntKey dec value of an integer
func (s *KVStorage) DecIntKey(key string, value int64) error {
	val, err := s.GetIntKey(key)
	if err != nil {
		return err
	}
	total := val + value
	return redisClient.Set(key, strconv.Itoa(int(total)), 0).Err()
}

// SetIntKey set int value of a key
func (s *KVStorage) SetIntKey(key string, value int64, exp time.Duration) error {
	return redisClient.Set(key, value, exp).Err()
}

// GetKey get int value of a key
func (s *KVStorage) GetIntKey(key string) (int64, error) {
	var value int64
	err := getInt64(key, &value)
	return value, err
}

// SetString key value of a key
func (s *KVStorage) SetString(key string, value string, exp time.Duration) error {
	return redisClient.Set(key, value, exp).Err()
}

// GetString get value of a key
func (s *KVStorage) GetString(key string) (string, error) {
	var value string
	err := getString(key, &value)
	return value, err
}

// RemoveKeys remove keys
func (s *KVStorage) RemoveKeys(keys ...string) error {
	_, err := redisClient.Del(keys...).Result()
	return err
}