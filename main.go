package main

import (
	"fmt"
	nats "github.com/nats-io/nats.go"
)

func main() {
	fmt.Printf("nats %s\n", nats.Version)
}