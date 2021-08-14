package rediscloudflarekv

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

// RedisKV is a type that implements a Redis key-value store
type RedisKV interface {
	RedisRead(key string) ([]byte, error)
	RedisWrite(key string, value []byte) (bool, error)
	RedisListKeysByPrefix(prefix string) ([]string, error)
	RedisDelete(key string) (bool, error)
}

// NewRedisClient returns a client to the Redis Server
func NewRedisClient(redisURL string) *KVOptions {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}
	redisClient := redis.NewClient(opt)
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln(err)
	}
	kvOptions := KVOptions{
		redisClient: redisClient,
	}
	return &kvOptions
}

// RedisRead returns value for that key
func (opt *KVOptions) RedisRead(key string) ([]byte, error) {
	value, err := opt.redisClient.Get(context.Background(), key).Bytes()
	if err != nil {
		return nil, err
	}
	return value, nil
}

// RedisWrite writes a value identified by a key.
func (opt *KVOptions) RedisWrite(key string, value []byte) (bool, error) {
	if err := opt.redisClient.Set(context.Background(), key, value, 0).Err(); err != nil {
		return false, err
	}
	return true, nil
}

// RedisListKeysByPrefix returns keys that matches the prefix
func (opt *KVOptions) RedisListKeysByPrefix(prefix string) ([]string, error) {
	var keys []string
	iter := opt.redisClient.Scan(context.Background(), 0, prefix+"*", 0).Iterator()
	for iter.Next(context.Background()) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return keys, nil
}

// RedisDelete deletes a key and value
func (opt *KVOptions) RedisDelete(key string) (bool, error) {
	if err := opt.redisClient.Del(context.Background(), key).Err(); err != nil {
		return false, err
	}
	return true, nil
}
