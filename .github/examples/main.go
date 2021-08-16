package main

import (
	"fmt"
	"log"
	"os"

	rediscloudflarekv "github.com/dineshsonachalam/redis-cloudflare-kv"
)

func main() {
	// Construct a new KV Client object
	kvClient := rediscloudflarekv.New(
		// REDIS_URL -> TCP Connection:  redis://<user>:<password>@<host>:<port>/<db_number>
		//              UNIX Connection: unix://<user>:<password>@</path/to/redis.sock>?db=<db_number>
		os.Getenv("REDIS_URL"),
		os.Getenv("CLOUDFLARE_ACCESS_KEY"),
		os.Getenv("CLOUDFLARE_EMAIL_ADDRESS"),
		os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
	)
	namespaceID := os.Getenv("TEST_NAMESPACE_ID")
	key1 := "dineshsonachalam.app1.aws.com"
	value1 := "gitlab.openai.com"
	key2 := "dineshsonachalam.app2.aws.com"
	value2 := "jira.openai.com"
	key3 := "dineshsonachalam.app3.aws.com"
	value3 := "confluence.openai.com"
	key4 := "dineshsonachalam.app4.aws.com"
	value4 := "grafana.openai.com"

	// Write writes a value identified by a key in Redis and Cloudflare KV
	status, err := kvClient.Write(key1, []byte(value1), namespaceID)
	if !status && err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Write operation is successful for key: %v\n", key1)

	// Read returns the value associated with the given key
	// If the key is not available in the Redis server,
	// it searches for the key in Cloudflare KV and if the key is available, it writes the key/value in the Redis.
	// Then it returns the value associated with the given key.
	value, err := kvClient.Read(key1, namespaceID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read operation is successful. Key: %v, Value: %v\n", key1, string(value))

	// Delete deletes a key/value in Redis and Cloudflare KV
	status, err = kvClient.Delete(key1, namespaceID)
	if !status && err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Delete operation is successful for key: %v\n", key1)

	kvClient.Write(key2, []byte(value2), namespaceID)
	kvClient.Write(key3, []byte(value3), namespaceID)
	kvClient.Write(key4, []byte(value4), namespaceID)

	// ListKeysByPrefix returns keys that matches the prefix
	// If there are no key's that matches the prefix in the Redis
	// We search for the prefix pattern in Cloudflare KV, if there are keys
	// that matches the prefix, we return the keys.
	keys, err := kvClient.ListKeysByPrefix("dinesh", namespaceID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ListKeysByPrefix operation is successful. Keys: %v\n", keys)
}

// Output:
// Write operation is successful for key: dineshsonachalam.app1.aws.com
// Read operation is successful. Key: dineshsonachalam.app1.aws.com, Value: gitlab.openai.com
// Delete operation is successful for key: dineshsonachalam.app1.aws.com
// ListKeysByPrefix operation is successful. Keys: [dineshsonachalam.app3.aws.com dineshsonachalam.app4.aws.com dineshsonachalam.app2.aws.com]
