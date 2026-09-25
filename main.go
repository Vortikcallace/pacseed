package main

import (
	"log/slog"
	"os"
	"pacseed/cli"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	cli.Execute()
}
