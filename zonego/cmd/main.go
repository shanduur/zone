package main

import (
	"log/slog"

	"github.com/shanduur/zone/zonego"
)

func main() {
	err := zonego.Zlogin("zone", "utility")
	if err != nil {
		slog.Error("zlogin failed", "error", err)
	}
}
