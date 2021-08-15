package rediscloudflarekv

import (
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var kvClient = New(
	os.Getenv("REDIS_URL"),
	os.Getenv("CLOUDFLARE_ACCESS_KEY"),
	os.Getenv("CLOUDFLARE_EMAIL_ADDRESS"),
	os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
)

func Test_New(t *testing.T) {
	assert.Equal(t, true, (New(
		os.Getenv("REDIS_URL"),
		os.Getenv("CLOUDFLARE_ACCESS_KEY"),
		os.Getenv("CLOUDFLARE_EMAIL_ADDRESS"),
		os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
	)) != nil)
}

func TestKV_Write(t *testing.T) {
	res, err := kvClient.Write("dineshsonachalam.app1.aws.com", []byte("gitlab.openai.com"), os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.Write("dineshsonachalam.app2.aws.com", []byte("jira.openai.com"), os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.Write("dineshsonachalam.app3.aws.com", []byte("confluence.openai.com"), os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.Write("dineshsonachalam.app4.aws.com", []byte("grafana.openai.com"), os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
}

func TestKV_Read(t *testing.T) {
	value, err := kvClient.Read("dineshsonachalam.app1.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, value, []byte("gitlab.openai.com"))
	}
	value, err = kvClient.Read("dineshsonachalam.app2.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, value, []byte("jira.openai.com"))
	}
	value, err = kvClient.Read("dineshsonachalam.app3.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, value, []byte("confluence.openai.com"))
	}
	value, err = kvClient.Read("dineshsonachalam.app4.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, value, []byte("grafana.openai.com"))
	}
	value, err = kvClient.Read("dineshsonachalam.app5.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.Error(t, err) {
		assert.Equal(t, len(value), 0)
	}
}

func TestKV_ListKeysByPrefix(t *testing.T) {
	keys, err := kvClient.ListKeysByPrefix("dinesh", os.Getenv("TEST_NAMESPACE_ID"))
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

func TestKV_Delete(t *testing.T) {
	res, err := kvClient.Delete("dineshsonachalam.app1.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.Delete("dineshsonachalam.app2.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.Delete("dineshsonachalam.app3.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.Delete("dineshsonachalam.app4.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	keys, err := kvClient.ListKeysByPrefix("dinesh", os.Getenv("TEST_NAMESPACE_ID"))
	expected_keys := []string(nil)
	if assert.NoError(t, err) {
		assert.Equal(t, keys, expected_keys)
	}
}
