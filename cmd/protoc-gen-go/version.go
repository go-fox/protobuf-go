package main

import "github.com/go-fox/protobuf-go/internal/version"

// Version return version
func Version() string {
	return version.String()
}
