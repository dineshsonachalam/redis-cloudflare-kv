<h1 align="center">
  <a href="https://github.com/marketplace/actions/markdown-autodocs">
    <img src="https://i.imgur.com/9Y1vNiT.png"/>
  </a>
</h1>
<p align="center">Go library for the Redis and Cloudflare KV</p>

<p align="center">
    <a href="https://sonarcloud.io/dashboard?id=redis-cloudflare-kv">
        <img src="https://sonarcloud.io/api/project_badges/quality_gate?project=redis-cloudflare-kv"/>
    </a>
</p>

<p align="center">
    <a href="https://goreportcard.com/report/github.com/dineshsonachalam/redis-cloudflare-kv">
        <img src="https://goreportcard.com/badge/github.com/dineshsonachalam/redis-cloudflare-kv">
    </a>
    <a href="https://github.com/dineshsonachalam/redis-cloudflare-kv/actions/workflows/tests.yml">
        <img src="https://github.com/dineshsonachalam/redis-cloudflare-kv/actions/workflows/tests.yml/badge.svg"/>
    </a>
    <a href="https://pkg.go.dev/github.com/dineshsonachalam/redis-cloudflare-kv">
        <img src="https://pkg.go.dev/badge/github.com/dineshsonachalam/redis-cloudflare-kv.svg" alt="Go Reference">
    </a>
</p>

Ask questions in the <a href ="https://github.com/dineshsonachalam/redis-cloudflare-kv/issues">GitHub issues</a>

## Installation

You need a working Go environment.

```
go get github.com/dineshsonachalam/redis-cloudflare-kv
```

## Quickstart

<!-- MARKDOWN-AUTO-DOCS:START (CODE:src=./.github/examples/main.go) -->
<!-- The below code snippet is automatically added from ./.github/examples/main.go -->
```go
package main

import (
	"fmt"
	"log"
	"os"

	rediscloudflarekv "github.com/dineshsonachalam/redis-cloudflare-kv"
)

func main() {
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

	status, err := kvClient.Write(key1, []byte(value1), namespaceID)
	if !status && err != nil {
		log.Fatalln("Write operation failed. Status: %v, Error: %v", err)
	}
	fmt.Printf("Write operation is successful for key: %v\n", key1)
	value, err := kvClient.Read(key1, namespaceID)
	if err != nil {
		log.Fatalln("Read operation failed, Error: %v", err)
	}
	fmt.Printf("Read operation is successful. Key: %v, Value: %v\n", key1, string(value))
	status, err = kvClient.Delete(key1, namespaceID)
	if !status && err != nil {
		log.Fatalln("Delete operation failed. Status: %v, Error: %v", status, err)
	}
	fmt.Printf("Delete operation is successful for key: %v\n", key1)

	kvClient.Write(key2, []byte(value2), namespaceID)
	kvClient.Write(key3, []byte(value3), namespaceID)
	kvClient.Write(key4, []byte(value4), namespaceID)

	keys, err := kvClient.ListKeysByPrefix("dinesh", namespaceID)
	if err != nil {
		log.Fatalln("ListKeysByPrefix operation failed, Err: %v", err)
	}
	fmt.Printf("ListKeysByPrefix operation is successful. Keys: %v\n", keys)
}

// Output:
// Write operation is successful for key: dineshsonachalam.app1.aws.com
// Read operation is successful. Key: dineshsonachalam.app1.aws.com, Value: gitlab.openai.com
// Delete operation is successful for key: dineshsonachalam.app1.aws.com
// ListKeysByPrefix operation is successful. Keys: [dineshsonachalam.app3.aws.com dineshsonachalam.app4.aws.com dineshsonachalam.app2.aws.com]
```
<!-- MARKDOWN-AUTO-DOCS:END -->