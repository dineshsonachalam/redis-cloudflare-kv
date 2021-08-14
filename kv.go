package rediscloudflarekv

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/go-redis/redis/v8"
)

// KVOptions represent Cloudflare and Redis client information
type KVOptions struct {
	redisClient      *redis.Client
	cloudflareClient *cloudflare.API
}

// KV is a type that implements a Redis and Cloudflare key-value store.
type KV interface {
	Read(key string, namespaceID string) ([]byte, error)
	Write(key string, value []byte, namespaceID string) (bool, error)
	ListKeysByPrefix(prefix string, namespaceID string) ([]string, error)
	Delete(key string, namespaceID string) (bool, error)
	RedisRead(key string) ([]byte, error)
	RedisWrite(key string, value []byte) (bool, error)
	RedisListKeysByPrefix(prefix string) ([]string, error)
	RedisDelete(key string) (bool, error)
	CloudflareKVRead(key string, namespaceID string) ([]byte, error)
	CloudflareKVWrite(key string, value []byte, namespaceID string) (bool, error)
	CloudflareKVListKeysByPrefix(prefix string, namespaceID string) ([]string, error)
	CloudflareKVDelete(key string, namespaceID string) (bool, error)
}

// New returns a client for Redis and CloudFlare KV
func New(redisURL string, cloudflareApiKey string, cloudflareEmail string, cloudflareAccountID string) *KVOptions {
	kvOptions := KVOptions{
		redisClient:      NewRedisClient(redisURL),
		cloudflareClient: NewCloudflareClient(cloudflareApiKey, cloudflareEmail, cloudflareAccountID),
	}
	return &kvOptions
}

// Read returns the value associated with the given key in the given redis server or Cloudflare KV namespace
func (opt *KVOptions) Read(key string, namespaceID string) ([]byte, error) {
	value, err := opt.RedisRead(key)
	if err != nil {
		value, err = opt.CloudflareKVRead(key, namespaceID)
		if err != nil {
			return nil, err
		}
		opt.RedisWrite(key, value)
	}
	return value, nil
}

// Write writes a value identified by a key.
func (opt *KVOptions) Write(key string, value []byte, namespaceID string) (bool, error) {
	resp, err := opt.RedisWrite(key, value)
	if err != nil {
		return false, err
	} else if !resp {
		return false, nil
	}
	resp, err = opt.CloudflareKVWrite(key, value, namespaceID)
	if err != nil {
		return false, err
	} else if !resp {
		return false, nil
	}
	return true, nil
}

// ListKeysByPrefix returns keys that matches the prefix
func (opt *KVOptions) ListKeysByPrefix(prefix string, namespaceID string) ([]string, error) {
	keys, err := opt.RedisListKeysByPrefix(prefix)
	if err != nil {
		return nil, err
	} else if len(keys) > 0 {
		return keys, nil
	} else {
		keys, err = opt.CloudflareKVListKeysByPrefix(prefix, namespaceID)
		if err != nil {
			return nil, err
		}
	}
	return keys, nil
}

// Delete deletes a key and value
func (opt *KVOptions) Delete(key string, namespaceID string) (bool, error) {
	resp, err := opt.RedisDelete(key)
	if err != nil {
		return false, err
	} else if !resp {
		return false, nil
	}
	resp, err = opt.CloudflareKVDelete(key, namespaceID)
	if err != nil {
		return false, err
	} else if !resp {
		return false, nil
	}
	return true, nil
}
