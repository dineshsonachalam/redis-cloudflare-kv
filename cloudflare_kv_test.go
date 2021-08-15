package rediscloudflarekv

import (
	"os"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCloudflareKV_NewCloudflareClient(t *testing.T) {
	cloudflareClient := NewCloudflareClient(
		os.Getenv("CLOUDFLARE_ACCESS_KEY"),
		os.Getenv("CLOUDFLARE_EMAIL_ADDRESS"),
		os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
	)
	assert.Equal(t, true, cloudflareClient != nil)
	cloudflareClient = NewCloudflareClient(
		"test",
		"test",
		"test",
	)
	assert.Equal(t, false, cloudflareClient == nil)
}

func TestCloudflareKV_CloudflareKVWrite(t *testing.T) {
	res, err := kvClient.CloudflareKVWrite("dineshsonachalam.app1.aws.com", []byte("gitlab.openai.com"), os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.CloudflareKVWrite("dineshsonachalam.app2.aws.com", []byte("jira.openai.com"), os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.CloudflareKVWrite("dineshsonachalam.app3.aws.com", []byte("confluence.openai.com"), os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.CloudflareKVWrite("dineshsonachalam.app4.aws.com", []byte("grafana.openai.com"), os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.CloudflareKVWrite("dineshsonachalam.app5.aws.com", []byte("swagger.openai.com"), "test")
	if assert.Error(t, err) {
		assert.Equal(t, false, res)
	}
}

func TestCloudflareKV_CloudflareKVRead(t *testing.T) {
	time.Sleep(60 * time.Second)

	value, err := kvClient.CloudflareKVRead("dineshsonachalam.app1.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, value, []byte("gitlab.openai.com"))
	}
	value, err = kvClient.CloudflareKVRead("dineshsonachalam.app2.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, value, []byte("jira.openai.com"))
	}
	value, err = kvClient.CloudflareKVRead("dineshsonachalam.app3.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, value, []byte("confluence.openai.com"))
	}
	value, err = kvClient.CloudflareKVRead("dineshsonachalam.app4.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, value, []byte("grafana.openai.com"))
	}
	value, err = kvClient.CloudflareKVRead("dineshsonachalam.app5.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.Error(t, err) {
		assert.Equal(t, len(value), 0)
	}
	value, err = kvClient.CloudflareKVRead("dineshsonachalam.app1.aws.com", "test")
	if assert.Error(t, err) {
		assert.Equal(t, value, []byte(nil))
	}
}

func TestCloudflareKV_CloudflareKVListKeysByPrefix(t *testing.T) {
	keys, err := kvClient.CloudflareKVListKeysByPrefix("dinesh", os.Getenv("TEST_NAMESPACE_ID"))
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
	keys, err = kvClient.CloudflareKVListKeysByPrefix("dinesh", "test")
	if assert.Error(t, err) {
		assert.Equal(t, keys, []string(nil))
	}
}

func TestCloudflareKV_CloudflareKVDelete(t *testing.T) {
	res, err := kvClient.CloudflareKVDelete("dineshsonachalam.app1.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.CloudflareKVDelete("dineshsonachalam.app2.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.CloudflareKVDelete("dineshsonachalam.app3.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.CloudflareKVDelete("dineshsonachalam.app4.aws.com", os.Getenv("TEST_NAMESPACE_ID"))
	if assert.NoError(t, err) {
		assert.Equal(t, true, res)
	}
	res, err = kvClient.CloudflareKVDelete("dineshsonachalam.app4.aws.com", "test")
	if assert.Error(t, err) {
		assert.Equal(t, res, false)
	}
	keys, err := kvClient.CloudflareKVListKeysByPrefix("dinesh", os.Getenv("TEST_NAMESPACE_ID"))
	expected_keys := []string(nil)
	if assert.NoError(t, err) {
		assert.Equal(t, keys, expected_keys)
	}
}
