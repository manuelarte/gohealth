//go:build tools

package main

import (
	_ "github.com/boumenot/gocover-cobertura"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "gotest.tools/gotestsum"
)
