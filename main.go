package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/miscord-dev/kubevctx/pkg/cmd"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := cmd.RootCommand().ExecuteContext(ctx)

	if err != nil {
		slog.Error("kubectx failed", "error", err)
		os.Exit(1)
	}
}
