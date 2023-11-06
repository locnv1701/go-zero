package redis

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
)

func init() {
	// Remote Cache Configuration Value

	redisCfg.Host = "localhost"
	redisCfg.Port = "6379"
	redisCfg.Password = ""
	redisCfg.Name = 1

	RedisCache = redisConnect()
}

const (
	_remoteKeyNotExist = redis.Nil
)

// Redis Configuration Struct
type redisConfig struct {
	Host     string
	Port     string
	Password string
	Name     int
}

var redisCfg = &redisConfig{}

type RedisCacheStore struct {
	RWMutex sync.RWMutex
	Limiter *redis_rate.Limiter
	Remote  *redis.Client
}

// Redis Cache Variable
var RedisCache *RedisCacheStore

// Redis Connect Function
func redisConnect() *RedisCacheStore {
	// Initialize Connection
	remote := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Host + ":" + redisCfg.Port,
		Password: redisCfg.Password,
		DB:       redisCfg.Name,
	})

	// Test remote connection
	_, err := remote.Ping().Result()
	if err != nil {
		fmt.Println("REDIS ERROR", err)
	} else {
		fmt.Println("REDIS OK")
	}

	// Return Connection
	return &RedisCacheStore{
		Limiter: redis_rate.NewLimiter(remote),
		Remote:  remote,
	}
}

// Get method to check redis server connection
func (redisCache *RedisCacheStore) Ping() (string, error) {
	return redisCache.Remote.Ping().Result()
}

// Set method to set cache by given key with time to live
func (redisCache *RedisCacheStore) Set(key string, value any, timeToLive time.Duration) error {
	redisCache.RWMutex.Lock()
	defer redisCache.RWMutex.Unlock()

	byteValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// Set to redis cache
	return redisCache.Remote.Set(key, byteValue, timeToLive).Err()
}

// Get method to retrieve the value of a key. If not present, returns false.
func (redisCache *RedisCacheStore) Get(key string, timeToLive time.Duration) ([]byte, bool, error) {
	// Get redis cache
	byteValue, err := redisCache.Remote.Get(key).Bytes()
	if err == _remoteKeyNotExist {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}

	return byteValue, true, nil

}

// Invalidate  method to delete a key from cahce.
func (redisCache *RedisCacheStore) Invalidate(key string) error {
	redisCache.RWMutex.Lock()
	defer redisCache.RWMutex.Unlock()

	// Invalidate local cache

	return redisCache.Remote.Del(key).Err()
}

// Close method clear and then close the cache Store.
func (redisCache *RedisCacheStore) Close() {
	redisCache.Remote.Close()
}
