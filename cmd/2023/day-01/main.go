package main

import (
	"log/slog"
	"os"
)

func main() {
	if err := process(os.Args); err != nil {
		slog.Error("process", "err", err)
		os.Exit(1)
	}
	os.Exit(0)
}
