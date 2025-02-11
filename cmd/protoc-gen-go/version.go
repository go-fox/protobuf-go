package main

import "google.golang.org/protobuf/internal/version"

// Version return version
func Version() string {
	return version.String()
}
