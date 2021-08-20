package main

import (
	"fmt"
	"log"
	"os"
	"time"

	rediscloudflarekv "github.com/dineshsonachalam/redis-cloudflare-kv"
)

var kvClient = rediscloudflarekv.New(
	os.Getenv("REDIS_URL"),
	os.Getenv("CLOUDFLARE_ACCESS_KEY"),
	os.Getenv("CLOUDFLARE_EMAIL_ADDRESS"),
	os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
)

func measureTime(funcName string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time taken by %s function is %v \n", funcName, time.Since(start))
		fmt.Println("==============================================================")
	}
}

func RedisRead(key string) {
	defer measureTime("RedisRead")()
	value, err := kvClient.RedisRead(key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v : %v \n", key, string(value))
}

func CloudflareKVRead(key string, namespaceID string) {
	defer measureTime("CloudflareKVRead")()
	value, err := kvClient.CloudflareKVRead(key, namespaceID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v : %v \n", key, string(value))
}

func main() {
	key := "opensource.facebook.react-native"
	namespaceID := os.Getenv("TEST_NAMESPACE_ID")
	RedisRead(key)
	CloudflareKVRead(key, namespaceID)
}

// dineshsonachalam@macbook examples % go run latency_check.go
// opensource.facebook.react-native : A framework for building native apps with React.
// Time taken by RedisRead function is 124.216708ms
// ==============================================================
// opensource.facebook.react-native : A framework for building native apps with React.
// Time taken by CloudflareKVRead function is 1.604654375s
// ==============================================================
// dineshsonachalam@macbook examples %
