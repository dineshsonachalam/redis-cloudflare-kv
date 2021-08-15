package rediscloudflarekv

import (
	"context"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedis_NewRedisClient(t *testing.T) {
	redisClient := NewRedisClient(os.Getenv("REDIS_URL"))
	res, err := redisClient.Ping(context.Background()).Result()
	if assert.NoError(t, err) {
		assert.Equal(t, "PONG", res)
	}
}

func TestRedis_RedisWrite(t *testing.T) {
	res, err := kvClient.RedisWrite("dineshsonachalam.app1.aws.com", []byte("gitlab.openai.com"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.RedisWrite("dineshsonachalam.app2.aws.com", []byte("jira.openai.com"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.RedisWrite("dineshsonachalam.app3.aws.com", []byte("confluence.openai.com"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.RedisWrite("dineshsonachalam.app4.aws.com", []byte("grafana.openai.com"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
}

func TestRedis_RedisRead(t *testing.T) {
	value, err := kvClient.RedisRead("dineshsonachalam.app1.aws.com")
	if assert.NoError(t, err) {
		assert.Equal(t, value, []byte("gitlab.openai.com"))
	}
	value, err = kvClient.RedisRead("dineshsonachalam.app2.aws.com")
	if assert.NoError(t, err) {
		assert.Equal(t, value, []byte("jira.openai.com"))
	}
	value, err = kvClient.RedisRead("dineshsonachalam.app3.aws.com")
	if assert.NoError(t, err) {
		assert.Equal(t, value, []byte("confluence.openai.com"))
	}
	value, err = kvClient.RedisRead("dineshsonachalam.app4.aws.com")
	if assert.NoError(t, err) {
		assert.Equal(t, value, []byte("grafana.openai.com"))
	}
	value, err = kvClient.RedisRead("dineshsonachalam.app5.aws.com")
	if assert.Error(t, err) {
		assert.Equal(t, len(value), 0)
	}
}

func TestRedis_RedisListKeysByPrefix(t *testing.T) {
	keys, err := kvClient.RedisListKeysByPrefix("dinesh")
	expected_keys := []string{
		"dineshsonachalam.app1.aws.com",
		"dineshsonachalam.app2.aws.com",
		"dineshsonachalam.app3.aws.com",
		"dineshsonachalam.app4.aws.com",
	}
	sort.Strings(keys)
	if assert.NoError(t, err) {
		assert.Equal(t, keys, expected_keys)
	}
}

func TestRedis_RedisDelete(t *testing.T) {
	res, err := kvClient.RedisDelete("dineshsonachalam.app1.aws.com")
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.RedisDelete("dineshsonachalam.app2.aws.com")
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.RedisDelete("dineshsonachalam.app3.aws.com")
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.RedisDelete("dineshsonachalam.app4.aws.com")
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	keys, err := kvClient.RedisListKeysByPrefix("dinesh")
	expected_keys := []string(nil)
	if assert.NoError(t, err) {
		assert.Equal(t, keys, expected_keys)
	}
}
