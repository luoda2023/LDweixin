//go:build windows

package main

import (
	"context"

	"github.com/luoda2023/LDweixin/config"
)

func runRunAsUserStartupChecks(_ context.Context, _ *config.Config) error {
	return nil
}
