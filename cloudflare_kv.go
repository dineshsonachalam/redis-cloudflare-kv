package rediscloudflarekv

import (
	"context"
	"log"

	"github.com/cloudflare/cloudflare-go"
)

// NewCloudflareClient returns a new Cloudflare v4 API client
func NewCloudflareClient(apiKey string, email string, accountID string) *cloudflare.API {
	cloudflareClient, err := cloudflare.New(apiKey, email, cloudflare.UsingAccount(accountID))
	if err != nil {
		log.Fatalln(err)
	}
	return cloudflareClient
}

// CloudflareKVRead returns the value associated with the given key in the given namespace
func (opt *Client) CloudflareKVRead(key string, namespaceID string) ([]byte, error) {
	value, err := opt.cloudflareClient.ReadWorkersKV(
		context.Background(),
		namespaceID,
		key,
	)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// CloudflareKVWrite writes a value identified by a key.
func (opt *Client) CloudflareKVWrite(key string, value []byte, namespaceID string) (bool, error) {
	response, err := opt.cloudflareClient.WriteWorkersKV(
		context.Background(),
		namespaceID,
		key,
		value,
	)
	if err != nil {
		return false, err
	} else if !response.Success {
		return false, nil
	}
	return true, nil
}

// CloudflareKVListKeysByPrefix returns keys that matches the prefix
func (opt *Client) CloudflareKVListKeysByPrefix(prefix string, namespaceID string) ([]string, error) {
	var keys []string
	limit := 500
	options := cloudflare.ListWorkersKVsOptions{
		Limit:  &limit,
		Prefix: &prefix,
	}
	resp, err := opt.cloudflareClient.ListWorkersKVsWithOptions(
		context.Background(),
		namespaceID,
		options)
	if err != nil {
		return nil, err
	}
	for _, value := range resp.Result {
		keys = append(keys, value.Name)
	}
	return keys, nil
}

// CloudflareKVDelete deletes a key and value for a provided storage namespace
func (opt *Client) CloudflareKVDelete(key string, namespaceID string) (bool, error) {
	response, err := opt.cloudflareClient.DeleteWorkersKV(
		context.Background(),
		namespaceID,
		key,
	)
	if err != nil {
		return false, err
	} else if !response.Success {
		return false, nil
	}
	return true, nil
}
